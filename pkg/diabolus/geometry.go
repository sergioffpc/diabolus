package diabolus

import "math"

type Bounds2 struct{ Min, Max Point2 }

type Normal3 struct{ X, Y, Z float64 }

func (n Normal3) Div(f float64) Normal3 { return Normal3.Mul(n, 1/f) }

func (n Normal3) Dot(m Normal3) float64 { return n.X*m.X + n.Y*m.Y + n.Z*m.Z }

func (n Normal3) Len() float64 { return math.Sqrt(Normal3.Dot(n, n)) }

func (n Normal3) Mul(f float64) Normal3 { return Normal3{n.X * f, n.Y * f, n.Z * f} }

func (n Normal3) Normalize() Normal3 { return Normal3.Div(n, n.Len()) }

type Point2 struct{ X, Y float64 }

type Point3 struct{ X, Y, Z float64 }

func (p Point3) Add(q Point3) Point3 { return Point3{p.X + q.X, p.Y + q.Y, p.Z + q.Z} }

func (p Point3) AddVector(v Vector3) Point3 { return Point3{p.X + v.X, p.Y + v.Y, p.Z + v.Z} }

func (p Point3) Distance(q Point3) float64 { return Point3.Sub(p, q).Len() }

func (p Point3) DistanceSq(q Point3) float64 { return Point3.Sub(p, q).LenSq() }

func (p *Point3) DivAssignFloat(f float64) { p.MulAssignFloat(f) }

func (p Point3) MulFloat(f float64) Point3 { return Point3{p.X * f, p.Y * f, p.Z * f} }

func (p *Point3) MulAssignFloat(f float64) {
	p.X *= f
	p.Y *= f
	p.Z *= f
}

func (p Point3) Neg() Point3 { return Point3.MulFloat(p, -1) }

func (p Point3) Sub(q Point3) Vector3 { return Vector3{p.X - q.X, p.Y - q.Y, p.Z - q.Z} }

func (p Point3) SubVector(v Vector3) Point3 { return Point3{p.X - v.X, p.Y - v.Y, p.Z - v.Z} }

type Ray struct {
	O    Point3
	D    Vector3
	TMax float64
}

func (r Ray) Position(t float64) Point3 { return Point3.AddVector(r.O, Vector3.MulFloat(r.D, t)) }

type Vector3 struct{ X, Y, Z float64 }

func (v Vector3) Add(u Vector3) Vector3 { return Vector3{v.X + u.X, v.Y + u.Y, v.Z + u.Z} }

func (v Vector3) DivFloat(f float64) Vector3 { return Vector3.MulFloat(v, 1/f) }

func (v Vector3) Dot(u Vector3) float64 { return v.X*u.X + v.Y*u.Y + v.Z*u.Z }

func (v Vector3) Len() float64 { return math.Sqrt(v.LenSq()) }

func (v Vector3) LenSq() float64 { return Vector3.Dot(v, v) }

func (v Vector3) MulFloat(f float64) Vector3 { return Vector3{v.X * f, v.Y * f, v.Z * f} }

func (v Vector3) Neg() Vector3 { return Vector3.MulFloat(v, -1) }

func (v Vector3) Normalize() Vector3 { return Vector3.DivFloat(v, v.Len()) }
