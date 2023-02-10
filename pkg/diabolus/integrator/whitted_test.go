package integrator_test

import (
	"math"
	"sergioffpc/diabolus/pkg/diabolus"
	"sergioffpc/diabolus/pkg/diabolus/integrator"
	"sergioffpc/diabolus/pkg/diabolus/light"
	"sergioffpc/diabolus/pkg/diabolus/material"
	"sergioffpc/diabolus/pkg/diabolus/sampler"
	"sergioffpc/diabolus/pkg/diabolus/shape"
	"testing"
)

func BenchmarkWhittedIntegrator(b *testing.B) {
	integrator := integrator.NewWhittedIntegrator(0)
	ray := diabolus.Ray{
		O:    diabolus.Point3{},
		D:    diabolus.Vector3{X: 0, Y: 0, Z: 1},
		TMax: math.MaxFloat64,
	}
	scene := diabolus.Scene{
		Geometries: []diabolus.GeometricPrimitive{
			diabolus.NewGeometricPrimitive(
				"sphere0",
				shape.NewSphereShape(),
				material.NewMatteMaterial(diabolus.Spectrum{R: 0, G: 0, B: 1}),
				diabolus.TranslateTransform(0, 0, 2),
			),
		},
		Lights: []diabolus.LightPrimitive{
			diabolus.NewLightPrimitive(
				"light0",
				light.NewPointLight(diabolus.Spectrum{R: 1, G: 1, B: 1}),
				diabolus.TranslateTransform(-1, -1, -1),
			),
		},
	}
	sampler := sampler.NewRandomSampler()
	for i := 0; i < b.N; i++ {
		integrator.Render(ray, scene, sampler)
	}
}
