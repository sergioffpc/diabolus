package integrator_test

import (
	"sergioffpc/diabolus/pkg/diabolus"
	"sergioffpc/diabolus/pkg/diabolus/integrator"
	"testing"
)

func BenchmarkWhittedIntegrator(b *testing.B) {
	integrator := integrator.NewWhittedIntegrator()
	for i := 0; i < b.N; i++ {
		integrator.Render(diabolus.Ray{}, diabolus.Scene{})
	}
}
