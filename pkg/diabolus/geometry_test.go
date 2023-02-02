package diabolus_test

import (
	"sergioffpc/diabolus/pkg/diabolus"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVector3Add(t *testing.T) {
	a := diabolus.Vector3{1, 2, 3}
	b := diabolus.Vector3{4, 5, 6}

	assert.Equal(t, diabolus.Vector3{5, 7, 9}, diabolus.Vector3.Add(a, b))
}
