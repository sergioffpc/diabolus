package diabolus

import "io"

type Film interface {
	AddSample(x, y int, s Spectrum)

	Bounds() Bounds2

	Write(w io.Writer) error
}
