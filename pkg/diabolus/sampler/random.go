package sampler

import (
	"math/rand"
	"sergioffpc/diabolus/pkg/diabolus"
)

type RandomSampler struct{}

func NewRandomSampler() RandomSampler { return RandomSampler{} }

func (s RandomSampler) Sample2D() diabolus.Point2 {
	return diabolus.Point2{
		X: rand.Float64(),
		Y: rand.Float64(),
	}
}
