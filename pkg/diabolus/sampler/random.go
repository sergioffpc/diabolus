package sampler

import (
	"math/rand"
	"sergioffpc/diabolus/pkg/diabolus"
)

type RandomSampler struct {
	SamplesCount int
}

func NewRandomSampler(samplesCount int) RandomSampler {
	return RandomSampler{SamplesCount: samplesCount}
}

func (s RandomSampler) GetSample2D() []diabolus.Point2 {
	us := make([]diabolus.Point2, 0, s.SamplesCount)
	for i := 0; i < s.SamplesCount; i++ {
		p := diabolus.Point2{
			X: rand.Float64(),
			Y: rand.Float64(),
		}
		us = append(us, p)
	}
	return us
}
