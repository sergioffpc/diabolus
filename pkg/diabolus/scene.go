package diabolus

import "math"

type Interaction struct {
	P      Point3
	N      Normal3
	Wo     Vector3
	T      float64
	Object *GeometricPrimitive
}

func (i Interaction) SampleLi(p LightPrimitive, u Point2) Spectrum {
	iL := TransformInteraction(i, p.WorldToLight)
	wiL, li := p.Light.SampleLi(iL, u)
	f := i.Object.Material.F(iL, wiL)
	theta := Vector3.Dot(wiL, Vector3(iL.N))
	return f.Mul(li).MulFloat(math.Abs(theta))
}

func (i Interaction) Le() Spectrum {
	return Spectrum{}
}

type GeometricPrimitive struct {
	Shape         Shape
	Material      Material
	ObjectToWorld Matrix44
	WorldToObject Matrix44
}

func (p *GeometricPrimitive) intersect(ray Ray) (ok bool, isect Interaction) {
	rO := TransformRay(ray, p.WorldToObject)
	if hit, pos, n, t := p.Shape.Intersect(rO); hit {
		ok = true
		isect = Interaction{
			P:      TransformPoint3(pos, p.ObjectToWorld),
			N:      TransformNormal3(n, p.ObjectToWorld).Normalize(),
			Wo:     ray.D.Neg(),
			T:      t,
			Object: p,
		}
	}
	return ok, isect
}

type LightPrimitive struct {
	Light        Light
	LightToWorld Matrix44
	WorldToLight Matrix44
}

type Scene struct {
	Geometries []GeometricPrimitive
	Lights     []LightPrimitive
}

func (s Scene) Intersect(ray Ray) (ok bool, nearest Interaction) {
	for _, p := range s.Geometries {
		if hit, isect := p.intersect(ray); hit {
			ok = true
			nearest = isect
			ray.TMax = isect.T
		}
	}
	return ok, nearest
}
