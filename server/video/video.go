package video

import (
	"data2video/video/column"
	"fmt"

	"github.com/fogleman/gg"
)

type Renderer struct {
}

func (r *Renderer) Render(dc *gg.Context) {

	data := []float64{12.3, 44, 50, 0, 1, 20.456, 25, 12.5}
	xLabels := []string{"A", "B", "C", "D", "E", "F"}
	b := column.New(dc, 1280, 720, 40, 20, 50.0, data, xLabels)

	steps := []float64{0, 0.1, 0.5, 1}

	for i, step := range steps {
		b.Render(step)
		dc.SavePNG(fmt.Sprintf("out%d.png", i))
	}
}
