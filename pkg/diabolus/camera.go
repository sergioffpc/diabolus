package diabolus

type Camera interface {
	GenerateRay(fP Point2) Ray
}
