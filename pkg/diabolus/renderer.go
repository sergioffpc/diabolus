package diabolus

import "io"

type Renderer interface {
	Render(scene Scene, w io.Writer) error
}
