package diabolus

import "io"

type Film interface {
	AddSample(x, y int, s Spectrum)

	GetBounds() Bounds2

	Write(w io.Writer) error
}
