package main

import (
	"log"
	"os"
	"sergioffpc/diabolus/pkg/diabolus"
	"sergioffpc/diabolus/pkg/diabolus/camera"
	"sergioffpc/diabolus/pkg/diabolus/film"
	"sergioffpc/diabolus/pkg/diabolus/integrator"
	"sergioffpc/diabolus/pkg/diabolus/light"
	"sergioffpc/diabolus/pkg/diabolus/material"
	"sergioffpc/diabolus/pkg/diabolus/renderer"
	"sergioffpc/diabolus/pkg/diabolus/sampler"
	"sergioffpc/diabolus/pkg/diabolus/shape"

	"github.com/schollz/progressbar/v3"
)

func main() {
	width, height, msaaCount := 1280, 720, 8

	out, err := os.OpenFile("image.png", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	pb := progressbar.Default(int64(width * height * msaaCount))
	r := renderer.ProgressRenderer{
		Progress: pb,
		Renderer: renderer.NewSubdivRenderer(
			integrator.ProgressIntegrator{
				Progress:   pb,
				Integrator: integrator.NewWhittedIntegrator(),
			},
			camera.NewPerspectiveCamera(width, height, sampler.NewRandomSampler(msaaCount)),
			film.NewImageFilm(width, height, msaaCount),
			1,
		),
	}

	s := diabolus.Scene{
		Geometries: []diabolus.GeometricPrimitive{
			{
				Shape:         shape.NewSphereShape(),
				Material:      material.NewMatteMaterial(diabolus.Spectrum{R: 0, G: 0, B: 10}),
				ObjectToWorld: diabolus.Translate(0, 0, 2),
				WorldToObject: diabolus.Translate(0, 0, -2),
			},
		},
		Lights: []diabolus.LightPrimitive{
			{
				Light:        light.NewPointLight(diabolus.Spectrum{R: 1, G: 1, B: 1}),
				LightToWorld: diabolus.Translate(0, 0, -1),
				WorldToLight: diabolus.Translate(0, 0, 1),
			},
		},
	}

	err = r.Render(s, out)
	if err != nil {
		log.Fatal(err)
	}
}
