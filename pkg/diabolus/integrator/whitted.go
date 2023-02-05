package integrator

import (
	"math"
	"sergioffpc/diabolus/pkg/diabolus"
)

type WhittedIntegrator struct{}

func NewWhittedIntegrator() WhittedIntegrator {
	return WhittedIntegrator{}
}

func (i WhittedIntegrator) Render(ray diabolus.Ray, scene diabolus.Scene, sampler diabolus.Sampler) diabolus.Spectrum {
	return i.li(ray, scene, sampler)
}

func (i WhittedIntegrator) li(ray diabolus.Ray, scene diabolus.Scene, sampler diabolus.Sampler) diabolus.Spectrum {
	var l diabolus.Spectrum
	if ok, isect := scene.Intersect(ray); ok {
		l.AddAssign(isect.Le())
		for _, p := range scene.Lights {
			lP := p
			u := sampler.Sample2D()
			wi, li := isect.SampleLi(lP, u)
			lR := diabolus.Ray{
				O:    isect.P,
				D:    wi,
				TMax: math.MaxFloat64,
			}
			if !scene.IntersectP(lR) {
				l.AddAssign(li)
			}
		}
	} else {
		l = diabolus.Spectrum{R: 0, G: 0, B: 1}
	}
	return l
}
