package camera

import (
	"math"
	"sergioffpc/diabolus/pkg/diabolus"
)

type PerspectiveCamera struct {
	Width  int
	Height int
}

func NewPerspectiveCamera(width, height int) PerspectiveCamera {
	return PerspectiveCamera{
		Width:  width,
		Height: height,
	}
}

func (c PerspectiveCamera) GenerateRay(fP diabolus.Point2) diabolus.Ray {
	dx := (-0.5 + (fP.X / float64(c.Width))) * (float64(c.Width) / float64(c.Height))
	dy := -0.5 + (fP.Y / float64(c.Height))
	return diabolus.Ray{
		O:    diabolus.Point3{},
		D:    diabolus.Vector3{X: dx, Y: dy, Z: 1}.Normalize(),
		TMax: math.MaxFloat64,
	}
}
