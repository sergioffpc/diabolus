package light

import "sergioffpc/diabolus/pkg/diabolus"

type PointLight struct {
	I diabolus.Spectrum
}

func NewPointLight(i diabolus.Spectrum) PointLight {
	return PointLight{I: i}
}

func (l PointLight) SampleLi(isect diabolus.Interaction, u diabolus.Point2) (diabolus.Vector3, float64, diabolus.Spectrum) {
	wi := diabolus.Point3.Sub(diabolus.Point3{}, isect.P).Normalize()
	pdf := 1.0
	return wi, pdf, l.I.DivFloat(diabolus.Point3.DistanceSq(isect.P, diabolus.Point3{}))
}
