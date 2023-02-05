package diabolus

import (
	"math"
)

type Interaction struct {
	P      Point3
	N      Normal3
	Wo     Vector3
	T      float64
	Object *GeometricPrimitive
}

func (i Interaction) SampleLi(p LightPrimitive, u Point2) (Vector3, Spectrum) {
	iL := i.Transform(p.WorldToLight)
	wiL, pdf, li := p.Light.SampleLi(iL, u)
	f := i.Object.Material.F(iL, wiL)
	theta := Vector3.Dot(wiL, Vector3(iL.N))
	return wiL.Transform(p.LightToWorld), f.Mul(li).MulFloat(math.Abs(theta)).DivFloat(pdf)
}

func (i Interaction) Le() Spectrum {
	return Spectrum{}
}

type GeometricPrimitive struct {
	Label         string
	Shape         Shape
	Material      Material
	ObjectToWorld Transform
	WorldToObject Transform
}

func NewGeometricPrimitive(label string, shape Shape, material Material, transform Transform) GeometricPrimitive {
	return GeometricPrimitive{
		Label:         label,
		Shape:         shape,
		Material:      material,
		ObjectToWorld: transform,
		WorldToObject: transform.Inverse(),
	}
}

func (p *GeometricPrimitive) intersect(ray Ray) (ok bool, isect Interaction) {
	rO := ray.Transform(p.WorldToObject)
	if hit, pO, nO, t := p.Shape.Intersect(rO); hit {
		ok = true
		isect = Interaction{
			P:      pO.Transform(p.ObjectToWorld),
			N:      nO.Transform(p.ObjectToWorld).Normalize(),
			Wo:     ray.D.Neg(),
			T:      t,
			Object: p,
		}
	}
	return ok, isect
}

func (p *GeometricPrimitive) intersectP(ray Ray) bool {
	rO := ray.Transform(p.WorldToObject)
	return p.Shape.IntersectP(rO)
}

type LightPrimitive struct {
	Label        string
	Light        Light
	LightToWorld Transform
	WorldToLight Transform
}

func NewLightPrimitive(label string, light Light, transform Transform) LightPrimitive {
	return LightPrimitive{
		Label:        label,
		Light:        light,
		LightToWorld: transform,
		WorldToLight: transform.Inverse(),
	}
}

type Scene struct {
	Geometries []GeometricPrimitive
	Lights     []LightPrimitive
}

func (s Scene) Intersect(ray Ray) (ok bool, nearest Interaction) {
	for _, p := range s.Geometries {
		gP := p
		if hit, isect := gP.intersect(ray); hit {
			ok = true
			nearest = isect
			ray.TMax = isect.T
		}
	}
	return ok, nearest
}

func (s Scene) IntersectP(ray Ray) bool {
	for _, p := range s.Geometries {
		gP := p
		if gP.intersectP(ray) {
			return true
		}
	}
	return false
}
