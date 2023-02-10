package integrator

import (
	"math"
	"sergioffpc/diabolus/pkg/diabolus"
)

type WhittedIntegrator struct {
	MaxDepth int
}

func NewWhittedIntegrator(maxDepth int) WhittedIntegrator {
	return WhittedIntegrator{MaxDepth: maxDepth}
}

func (i WhittedIntegrator) Render(ray diabolus.Ray, scene diabolus.Scene, sampler diabolus.Sampler) diabolus.Spectrum {
	return i.li(ray, scene, sampler, i.MaxDepth)
}

func (i WhittedIntegrator) li(ray diabolus.Ray, scene diabolus.Scene, sampler diabolus.Sampler, depth int) diabolus.Spectrum {
	var l diabolus.Spectrum
	if depth <= 0 {
		return l
	}

	if ok, isect := scene.Intersect(ray); ok {
		l.AddAssign(isect.Le())
		for _, p := range scene.Lights {
			lP := p
			wi, li := isect.SampleLi(lP, sampler.Sample2D())
			lR := diabolus.Ray{
				O:    isect.P,
				D:    wi,
				TMax: math.MaxFloat64,
			}
			if !scene.IntersectP(lR) {
				l.AddAssign(li)
			}
			l.AddAssign(i.fr(ray, scene, isect, sampler, depth))
		}
	}
	return l
}

func (i WhittedIntegrator) fr(ray diabolus.Ray, scene diabolus.Scene, isect diabolus.Interaction, sampler diabolus.Sampler, depth int) diabolus.Spectrum {
	wi, f := isect.SampleF(sampler.Sample2D())
	sR := diabolus.Ray{
		O:    isect.P,
		D:    wi,
		TMax: math.MaxFloat64,
	}
	return f.Mul(i.li(sR, scene, sampler, depth-1))
}
