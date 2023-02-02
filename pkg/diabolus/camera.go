package diabolus

type Camera interface {
	RaySamples(x, y int) []Ray
}
