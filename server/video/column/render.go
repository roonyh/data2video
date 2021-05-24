package column

import (
	"github.com/fogleman/gg"
)

type Column struct {
	*gg.Context
	width      int
	height     int
	margin     int
	gutter     int
	max        float64
	colWidth   float64
	colHeight  float64
	yAxisWidth float64
	data       []float64
}

func New(dc *gg.Context, width, height, margin, gutter int, max float64, data []float64, fields []string) *Column {
	yAxisWidth := 60.0

	colNum := len(data)
	colWidth := (float64(width-(2*margin)-(colNum+1)*gutter) - yAxisWidth) / float64(colNum)
	colHeight := float64(height - (2 * margin) - (2 * gutter))

	return &Column{
		Context:    dc,
		width:      width,
		height:     height,
		margin:     margin,
		gutter:     gutter,
		max:        max,
		colWidth:   colWidth,
		colHeight:  colHeight,
		yAxisWidth: yAxisWidth,
		data:       data,
	}
}

func (b *Column) Render(step float64) {
	b.SetRGBA(0, 0, 0, 0)
	b.Clear()
	b.SetRGBA(0, 0, 0.2, 0.2)
	b.DrawRectangle(
		0,
		0,
		float64(b.width),
		float64(b.height),
	)

	b.Fill()

	b.renderYAxis(step)
	b.renderXAxis(step)

	totalData := len(b.data)
	interval := 1.0 / float64(totalData)

	for i, d := range b.data {
		until := interval * float64(i)
		newStep := 0.0
		if step >= until {
			newStep = (step - until) / (1 - until)
		}
		b.renderColumn(d, i, newStep)
	}
}

func (b *Column) renderColumn(value float64, pos int, step float64) {
	b.SetRGBA(0.2, 0.8, 0.7, 0.8)

	height := (b.colHeight / b.max) * value * step
	x := float64(b.margin) + b.yAxisWidth + float64(b.gutter) + float64(pos)*(b.colWidth+float64(b.gutter))

	b.DrawRectangle(
		x,
		b.colHeight-height+float64(b.margin)+float64(b.gutter),
		b.colWidth,
		height,
	)

	b.Fill()
}

func (b *Column) renderYAxis(step float64) {
	b.SetRGBA(0, 1, 1, 0.6)
	b.SetLineWidth(6)
	b.DrawLine(
		float64(b.margin)+b.yAxisWidth,
		float64(b.height-b.margin),
		float64(b.margin)+b.yAxisWidth,
		float64(b.height-b.margin)-step*(float64(b.height-2*b.margin)),
	)
	b.Stroke()
}

func (b *Column) renderXAxis(step float64) {
	b.SetRGBA(0, 1, 1, 0.6)
	b.SetLineWidth(6)
	b.DrawLine(
		float64(b.margin)+b.yAxisWidth,
		float64(b.height)-float64(b.margin),
		float64(b.margin)+b.yAxisWidth+step*(float64(b.width)-float64(2*b.margin)-b.yAxisWidth),
		float64(b.height)-float64(b.margin),
	)
	b.Stroke()
}
