package diabolus

import "math"

type Material interface {
	F(isect Interaction, wi Vector3) Spectrum

	SampleF(isect Interaction, u Point2) (Vector3, float64, Spectrum)
}

type Bsdf struct {
	Bxdfs []Reflection
}

func (b Bsdf) F(isect Interaction, wi Vector3) Spectrum {
	var f Spectrum
	if Vector3.Dot(wi, Vector3(isect.N))*Vector3.Dot(isect.Wo, Vector3(isect.N)) > 0 {
		for _, bxdf := range b.Bxdfs {
			f.AddAssign(bxdf.F(isect.Wo, wi))
		}
	}
	return f
}

func (b Bsdf) SampleF(isect Interaction, u Point2) (Vector3, float64, Spectrum) {
	wi := ConcentricSampleHemisphere(u)
	if isect.Wo.Z < 0 {
		wi.Z *= -1
	}

	var pdf float64
	if isect.Wo.Z*wi.Z > 0 {
		pdf += math.Abs(wi.Z) * (1 / math.Pi)
	}

	f := b.F(isect, wi)

	return wi, pdf, f
}
