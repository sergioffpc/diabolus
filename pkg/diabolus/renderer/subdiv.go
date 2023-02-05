package renderer

import (
	"io"
	"math"
	"sergioffpc/diabolus/pkg/diabolus"
	"sync"
)

type SubdivRenderer struct {
	Integrator   diabolus.Integrator
	Camera       diabolus.Camera
	Film         diabolus.Film
	Sampler      diabolus.Sampler
	SamplesCount int
	SubdivCount  int
}

func NewSubdivRenderer(
	integrator diabolus.Integrator,
	camera diabolus.Camera,
	film diabolus.Film,
	sampler diabolus.Sampler,
	samplesCount int,
	subdivCount int,
) SubdivRenderer {
	return SubdivRenderer{
		Integrator:   integrator,
		Camera:       camera,
		Film:         film,
		Sampler:      sampler,
		SamplesCount: samplesCount,
		SubdivCount:  subdivCount,
	}
}

func (r SubdivRenderer) Render(scene diabolus.Scene, w io.Writer) error {
	var wg sync.WaitGroup
	for _, subdiv := range r.subdivs() {
		wg.Add(1)
		go func(subdiv diabolus.Bounds2) {
			defer wg.Done()
			r.render(scene, subdiv)
		}(subdiv)
	}
	wg.Wait()
	return r.Film.Write(w)
}

func (r SubdivRenderer) subdivs() []diabolus.Bounds2 {
	width, height := int(r.Film.Bounds().Max.X), int(r.Film.Bounds().Max.Y)
	size := int(math.Sqrt(float64(width*height) / float64(r.SubdivCount)))
	bounds := make([]diabolus.Bounds2, 0, (width*height)/(size*size)+1)
	for y := 0; y < height; y += size {
		for x := 0; x < width; x += size {
			bounds = append(bounds, diabolus.Bounds2{
				Min: diabolus.Point2{
					X: float64(x),
					Y: float64(y),
				},
				Max: diabolus.Point2{
					X: diabolus.Clamp(float64(x+size), 0, float64(width)),
					Y: diabolus.Clamp(float64(y+size), 0, float64(height)),
				},
			})
		}
	}
	return bounds
}

func (r SubdivRenderer) render(scene diabolus.Scene, subdiv diabolus.Bounds2) {
	for y := int(subdiv.Min.Y); y < int(subdiv.Max.Y); y++ {
		for x := int(subdiv.Min.X); x < int(subdiv.Max.X); x++ {
			p := diabolus.Point2{X: float64(x), Y: float64(y)}
			for i := 0; i < r.SamplesCount; i++ {
				u := r.Sampler.Sample2D()
				ray := r.Camera.GenerateRay(diabolus.Point2.Add(p, u))
				s := r.Integrator.Render(ray, scene, r.Sampler)
				r.Film.AddSample(x, y, s)
			}
		}
	}
}
