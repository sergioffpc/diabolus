package diabolus

type Shape interface {
	Intersect(ray Ray) (bool, Point3, Normal3, float64)

	IntersectP(ray Ray) bool
}
