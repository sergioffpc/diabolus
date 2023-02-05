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
	width, height := 1280, 720
	subdivsCount, samplesCount := 1, 1

	out, err := os.OpenFile("image.png", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	pb := progressbar.Default(int64(width * height * samplesCount))
	r := renderer.ProgressRenderer{
		Progress: pb,
		Renderer: renderer.NewSubdivRenderer(
			integrator.ProgressIntegrator{
				Progress:   pb,
				Integrator: integrator.NewWhittedIntegrator(),
			},
			camera.NewPerspectiveCamera(width, height),
			film.NewImageFilm(width, height, samplesCount),
			sampler.NewRandomSampler(),
			samplesCount,
			subdivsCount,
		),
	}

	s := diabolus.Scene{
		Geometries: []diabolus.GeometricPrimitive{
			diabolus.NewGeometricPrimitive(
				"sphere0",
				shape.NewSphereShape(),
				material.NewMatteMaterial(diabolus.Spectrum{R: 1, G: 0, B: 0}),
				diabolus.TranslateTransform(0, 0, 5),
			),
			diabolus.NewGeometricPrimitive(
				"sphere1",
				shape.NewSphereShape(),
				material.NewMatteMaterial(diabolus.Spectrum{R: 0, G: 1, B: 0}),
				diabolus.TranslateTransform(-1.5, 0, 2).Mul(diabolus.ScaleTransform(0.5, 0.5, 0.5)),
			),
		},
		Lights: []diabolus.LightPrimitive{
			diabolus.NewLightPrimitive(
				"light0",
				light.NewPointLight(diabolus.Spectrum{R: 50, G: 50, B: 50}),
				diabolus.TranslateTransform(-2, 0, 0),
			),
		},
	}

	err = r.Render(s, out)
	if err != nil {
		log.Fatal(err)
	}
}
