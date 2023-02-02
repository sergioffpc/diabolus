package diabolus

type Material interface {
	F(isect Interaction, wi Vector3) Spectrum
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
