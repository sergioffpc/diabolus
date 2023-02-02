package diabolus

type Reflection interface {
	F(wo Vector3, wi Vector3) Spectrum
}
