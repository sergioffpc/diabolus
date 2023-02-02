package renderer

import (
	"io"
	"math"
	"sergioffpc/diabolus/pkg/diabolus"
	"sync"
)

type SubdivRenderer struct {
	integrator  diabolus.Integrator
	camera      diabolus.Camera
	film        diabolus.Film
	subdivCount int
}

func NewSubdivRenderer(
	integrator diabolus.Integrator,
	camera diabolus.Camera,
	film diabolus.Film,
	subdivCount int,
) SubdivRenderer {
	return SubdivRenderer{
		integrator:  integrator,
		camera:      camera,
		film:        film,
		subdivCount: subdivCount,
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
	return r.film.Write(w)
}

func (r SubdivRenderer) subdivs() []diabolus.Bounds2 {
	width, height := int(r.film.GetBounds().Max.X), int(r.film.GetBounds().Max.Y)
	size := int(math.Sqrt(float64(width*height) / float64(r.subdivCount)))
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
			for _, ray := range r.camera.RaySamples(x, y) {
				s := r.integrator.Render(ray, scene)
				r.film.AddSample(x, y, s)
			}
		}
	}
}
