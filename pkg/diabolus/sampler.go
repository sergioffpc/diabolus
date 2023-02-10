package diabolus

import "math"

type Sampler interface {
	Sample2D() Point2
}

func UniformSampleDisk(u Point2) Point2 {
	r := math.Sqrt(u.X)
	theta := 2 * math.Pi * u.Y
	return Point2{X: r * math.Cos(theta), Y: r * math.Sin(theta)}
}

func ConcentricSampleDisk(u Point2) Point2 {
	uOffset := Point2.SubVector(u.MulFloat(2), Vector2{1, 1})
	if uOffset.X == 0 && uOffset.Y == 0 {
		return Point2{}
	}

	var r, theta float64
	if math.Abs(uOffset.X) > math.Abs(uOffset.Y) {
		r = uOffset.X
		theta = (math.Pi / 4) * (uOffset.Y / uOffset.X)
	} else {
		r = uOffset.Y
		theta = (math.Pi/2 - math.Pi/4) * (uOffset.X / uOffset.Y)
	}
	return Point2{X: math.Cos(theta), Y: math.Sin(theta)}.MulFloat(r)
}

func ConcentricSampleHemisphere(u Point2) Vector3 {
	d := ConcentricSampleDisk(u)
	z := math.Sqrt(math.Max(0, 1-d.X*d.X-d.Y*d.Y))
	return Vector3{X: d.X, Y: d.Y, Z: z}
}
