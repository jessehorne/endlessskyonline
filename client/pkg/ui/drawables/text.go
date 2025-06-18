package drawables

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"image/color"
)

type Text struct {
	X     float64
	Y     float64
	Data  string
	Font  *text.GoTextFace
	Color color.Color
}

func NewText(x float64, y float64, data string) *Text {
	return &Text{
		X:     x,
		Y:     y,
		Data:  data,
		Font:  nil,
		Color: color.White,
	}
}

func (t *Text) Draw(screen *ebiten.Image) {
	if t.Font == nil {
		return
	}

	opts := &text.DrawOptions{}
	opts.GeoM.Translate(t.X, t.Y)
	opts.ColorScale.ScaleWithColor(t.Color)

	text.Draw(screen, t.Data, t.Font, opts)
}
