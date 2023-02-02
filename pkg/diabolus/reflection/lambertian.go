package reflection

import (
	"math"
	"sergioffpc/diabolus/pkg/diabolus"
)

type LambertianReflection struct {
	R diabolus.Spectrum
}

func (r LambertianReflection) F(wo diabolus.Vector3, wi diabolus.Vector3) diabolus.Spectrum {
	return diabolus.Spectrum.MulFloat(r.R, 1/math.Pi)
}
