package integrator

import "sergioffpc/diabolus/pkg/diabolus"

type WhittedIntegrator struct{}

func NewWhittedIntegrator() WhittedIntegrator {
	return WhittedIntegrator{}
}

func (i WhittedIntegrator) Render(ray diabolus.Ray, scene diabolus.Scene) diabolus.Spectrum {
	return i.li(ray, scene)
}

func (i WhittedIntegrator) li(ray diabolus.Ray, scene diabolus.Scene) diabolus.Spectrum {
	var l diabolus.Spectrum
	if ok, isect := scene.Intersect(ray); ok {
		l.AddAssign(isect.Le())
		for _, p := range scene.Lights {
			l.AddAssign(isect.SampleLi(p, diabolus.Point2{}))
		}
	}
	return l
}
