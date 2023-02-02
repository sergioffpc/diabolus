package diabolus

type Light interface {
	SampleLi(isect Interaction, u Point2) (Vector3, Spectrum)
}
