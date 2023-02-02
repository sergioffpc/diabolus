package diabolus

type Integrator interface {
	Render(ray Ray, scene Scene) Spectrum
}
