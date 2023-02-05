package diabolus_test

import (
	"sergioffpc/diabolus/pkg/diabolus"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatrix44Adj(t *testing.T) {
	a := diabolus.Matrix44{
		1, 1, 1, -1,
		1, 1, -1, 1,
		1, -1, 1, 1,
		-1, 1, 1, 1,
	}
	b := diabolus.Matrix44{
		-4, -4, -4, 4,
		-4, -4, 4, -4,
		-4, 4, -4, -4,
		4, -4, -4, -4,
	}

	assert.Equal(t, b, a.Adj())
}

func TestMatrix44Det(t *testing.T) {
	a := diabolus.Matrix44{
		1, 1, 1, -1,
		1, 1, -1, 1,
		1, -1, 1, 1,
		-1, 1, 1, 1,
	}

	assert.Equal(t, -16.0, a.Det())
}

func TestMatrix44Inv(t *testing.T) {
	a := diabolus.Matrix44{
		1, 1, 1, -1,
		1, 1, -1, 1,
		1, -1, 1, 1,
		-1, 1, 1, 1,
	}
	b := diabolus.Matrix44{
		1.0 / 4.0, 1.0 / 4.0, 1.0 / 4.0, -1.0 / 4.0,
		1.0 / 4.0, 1.0 / 4.0, -1.0 / 4.0, 1.0 / 4.0,
		1.0 / 4.0, -1.0 / 4.0, 1.0 / 4.0, 1.0 / 4.0,
		-1.0 / 4.0, 1.0 / 4.0, 1.0 / 4.0, 1.0 / 4.0,
	}

	assert.Equal(t, b, a.Inverse())
}

func TestMatrix44Mul(t *testing.T) {
	a := diabolus.Matrix44{
		1, 1, 1, -1,
		1, 1, -1, 1,
		1, -1, 1, 1,
		-1, 1, 1, 1,
	}

	assert.Equal(t, diabolus.Identity(), diabolus.Matrix44.Mul(a, a.Inverse()))
}

func TestMatrix44Transpose(t *testing.T) {
	a := diabolus.Matrix44{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 0, 1, 2,
		3, 4, 5, 6,
	}
	b := diabolus.Matrix44{
		1, 5, 9, 3,
		2, 6, 0, 4,
		3, 7, 1, 5,
		4, 8, 2, 6,
	}

	assert.Equal(t, b, a.Transpose())
}

func TestScale(t *testing.T) {
	a := diabolus.Matrix44{
		1, 0, 0, 0,
		0, 2, 0, 0,
		0, 0, 3, 0,
		0, 0, 0, 1,
	}

	assert.Equal(t, a, diabolus.ScaleTransform(1, 2, 3).M)
}

func TestScaleInv(t *testing.T) {
	a := diabolus.Matrix44{
		1, 0, 0, 0,
		0, 1.0 / 2, 0, 0,
		0, 0, 1.0 / 3, 0,
		0, 0, 0, 1,
	}

	assert.Equal(t, a, diabolus.ScaleTransform(1, 2, 3).Inv)
}

func TestTranslate(t *testing.T) {
	a := diabolus.Matrix44{
		1, 0, 0, 1,
		0, 1, 0, 2,
		0, 0, 1, 3,
		0, 0, 0, 1,
	}

	assert.Equal(t, a, diabolus.TranslateTransform(1, 2, 3).M)
}

func TestTranslateInv(t *testing.T) {
	a := diabolus.Matrix44{
		1, 0, 0, -1,
		0, 1, 0, -2,
		0, 0, 1, -3,
		0, 0, 0, 1,
	}

	assert.Equal(t, a, diabolus.TranslateTransform(1, 2, 3).Inv)
}
