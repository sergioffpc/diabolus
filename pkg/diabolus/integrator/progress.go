package integrator

import (
	"sergioffpc/diabolus/pkg/diabolus"

	"github.com/schollz/progressbar/v3"
)

type ProgressIntegrator struct {
	Progress   *progressbar.ProgressBar
	Integrator diabolus.Integrator
}

func (i ProgressIntegrator) Render(ray diabolus.Ray, scene diabolus.Scene) diabolus.Spectrum {
	i.Progress.Add(1)
	return i.Integrator.Render(ray, scene)
}
