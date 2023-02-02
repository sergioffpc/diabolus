package renderer

import (
	"io"
	"sergioffpc/diabolus/pkg/diabolus"

	"github.com/schollz/progressbar/v3"
)

type ProgressRenderer struct {
	Progress *progressbar.ProgressBar
	Renderer diabolus.Renderer
}

func (r ProgressRenderer) Render(scene diabolus.Scene, w io.Writer) error {
	defer r.Progress.Close()
	return r.Renderer.Render(scene, w)
}
