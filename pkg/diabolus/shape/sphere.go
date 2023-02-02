package shape

import (
	"math"
	"sergioffpc/diabolus/pkg/diabolus"
)

type SphereShape struct{}

func NewSphereShape() SphereShape { return SphereShape{} }

func (s SphereShape) Intersect(ray diabolus.Ray) (bool, diabolus.Point3, diabolus.Normal3, float64) {
	a := diabolus.Vector3.Dot(ray.D, ray.D)
	b := 2 * diabolus.Vector3.Dot(ray.D, diabolus.Vector3(ray.O))
	c := diabolus.Vector3.Dot(diabolus.Vector3(ray.O), diabolus.Vector3(ray.O)) - 1

	switch ok, t0, t1 := diabolus.QuadraticSolver(a, b, c); {
	case ok && t0 > 0 && t0 < ray.TMax:
		p := ray.Position(t0)
		n := diabolus.Normal3(diabolus.Point3.Sub(p, diabolus.Point3{})).Normalize()
		return true, p, n, t0
	case ok && t1 > 0 && t1 < ray.TMax:
		p := ray.Position(t1)
		n := diabolus.Normal3(diabolus.Point3.Sub(p, diabolus.Point3{})).Normalize()
		return true, p, n, t1
	default:
		return false, diabolus.Point3{}, diabolus.Normal3{}, math.MaxFloat64
	}
}
