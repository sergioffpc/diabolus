package diabolus

import "math"

type Interaction struct {
	P      Point3
	N      Normal3
	Wo     Vector3
	T      float64
	Object *GeometricPrimitive
}

func (i Interaction) SampleF(u Point2) (wi Vector3, f Spectrum) {
	iL := i.Transform(i.Object.WorldToObject)
	wiL, pdf, fS := i.Object.Material.SampleF(iL, u)
	if pdf > 0 {
		theta := Vector3.Dot(wiL, Vector3(iL.N))
		wi = wiL.Transform(i.Object.ObjectToWorld)
		f = fS.MulFloat(math.Abs(theta)).DivFloat(pdf)
	}
	return wi, f
}

func (i Interaction) SampleLi(p LightPrimitive, u Point2) (wi Vector3, li Spectrum) {
	iL := i.Transform(p.WorldToLight)
	wiL, pdf, liS := p.Light.SampleLi(iL, u)
	if pdf > 0 {
		f := i.Object.Material.F(iL, wiL)
		theta := Vector3.Dot(wiL, Vector3(iL.N))
		wi = wiL.Transform(p.LightToWorld)
		li = f.Mul(liS).MulFloat(math.Abs(theta)).DivFloat(pdf)
	}
	return wi, li
}

func (i Interaction) Le() Spectrum {
	return Spectrum{}
}
