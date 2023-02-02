package camera

import (
	"math"
	"sergioffpc/diabolus/pkg/diabolus"
)

type PerspectiveCamera struct {
	Width   int
	Height  int
	Sampler diabolus.Sampler
}

func NewPerspectiveCamera(width, height int, sampler diabolus.Sampler) PerspectiveCamera {
	return PerspectiveCamera{
		Width:   width,
		Height:  height,
		Sampler: sampler,
	}
}

func (c PerspectiveCamera) RaySamples(x, y int) []diabolus.Ray {
	us := c.Sampler.GetSample2D()
	rs := make([]diabolus.Ray, 0, len(us))
	for _, u := range us {
		dx := (-0.5 + ((float64(x) + u.X) / float64(c.Width))) * (float64(c.Width) / float64(c.Height))
		dy := -0.5 + ((float64(y) + u.Y) / float64(c.Height))
		r := diabolus.Ray{
			O:    diabolus.Point3{},
			D:    diabolus.Vector3{X: dx, Y: dy, Z: 1}.Normalize(),
			TMax: math.MaxFloat64,
		}
		rs = append(rs, r)
	}
	return rs
}
