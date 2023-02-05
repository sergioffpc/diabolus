package film

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"sergioffpc/diabolus/pkg/diabolus"
)

type ImageFilm struct {
	Width        int
	Height       int
	SamplesCount int

	pixels []diabolus.Spectrum
}

func NewImageFilm(width, height, samplesCount int) ImageFilm {
	return ImageFilm{
		Width:        width,
		Height:       height,
		SamplesCount: samplesCount,

		pixels: make([]diabolus.Spectrum, width*height),
	}
}

func (f ImageFilm) AddSample(x, y int, s diabolus.Spectrum) { f.pixels[f.indexOf(x, y)].AddAssign(s) }

func (f ImageFilm) Bounds() diabolus.Bounds2 {
	return diabolus.Bounds2{
		Min: diabolus.Point2{},
		Max: diabolus.Point2{X: float64(f.Width), Y: float64(f.Height)},
	}
}

func (f ImageFilm) Write(w io.Writer) error {
	img := image.NewRGBA(image.Rect(0, 0, f.Width, f.Height))
	for y := 0; y < f.Height; y++ {
		for x := 0; x < f.Width; x++ {
			s := f.pixels[f.indexOf(x, y)].DivFloat(float64(f.SamplesCount)).Clamp(0, 1)
			img.Set(x, y, color.RGBA{
				R: uint8(s.R * 255),
				G: uint8(s.G * 255),
				B: uint8(s.B * 255),
				A: 255,
			})
		}
	}
	return png.Encode(w, img)
}

func (f ImageFilm) indexOf(x, y int) int { return y*f.Width + x }
