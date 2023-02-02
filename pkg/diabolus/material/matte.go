package material

import (
	"sergioffpc/diabolus/pkg/diabolus"
	"sergioffpc/diabolus/pkg/diabolus/reflection"
)

type MatteMaterial struct{ diabolus.Bsdf }

func NewMatteMaterial(kd diabolus.Spectrum) MatteMaterial {
	return MatteMaterial{
		diabolus.Bsdf{
			Bxdfs: []diabolus.Reflection{reflection.LambertianReflection{R: kd}},
		},
	}
}

func (m MatteMaterial) F(isect diabolus.Interaction, wi diabolus.Vector3) diabolus.Spectrum {
	return m.Bsdf.F(isect, wi)
}
