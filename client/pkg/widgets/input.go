package widgets

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/jessehorne/endlessskyonline/client/pkg/ui"
	"golang.org/x/image/font"
	"image/color"
)

const (
	defaultInputWidth  = 100
	defaultInputHeight = 32
)

type Input struct {
	X               float64
	Y               float64
	Width           int
	Height          int
	Data            string
	Font            *text.GoTextFace
	Color           color.Color
	BackgroundColor color.Color
	BorderColor     color.Color
	BlinkingEnabled bool
	BlinkSpeed      int // how many times cursor blinks a second
	BlinkTimer      float64
	BlinkingUp      bool // if true, adding numbers to alpha, if false, subtracting
	CursorColor     color.RGBA
	CursorOffsetX   float64
	CursorOffsetY   float64
	CursorWidth     int
	CursorHeight    int
	CursorX         float64
	CursorY         float64
}

func NewInput(x, y float64) *Input {
	return &Input{
		X:               x,
		Y:               y,
		Width:           defaultInputWidth,
		Height:          defaultInputHeight,
		Font:            ui.TextFaces["roboto-32"],
		Color:           color.White,
		BackgroundColor: color.Black,
		BorderColor:     color.White,
		BlinkingEnabled: true,
		BlinkSpeed:      1,
		BlinkTimer:      0,
		BlinkingUp:      true,
		CursorColor:     color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0},
		CursorOffsetX:   5,
		CursorOffsetY:   2,
		CursorWidth:     2,
		CursorHeight:    28,
		CursorX:         x + 5,
		CursorY:         y + 2,
	}
}

func (i *Input) Update() {
	// handle cursor blinking
	i.BlinkTimer += 1

	if i.BlinkTimer > float64(60/i.BlinkSpeed) {
		i.BlinkTimer = 0

		i.BlinkingUp = !i.BlinkingUp
	}

	var a uint8
	if i.BlinkingUp {
		a = i.CursorColor.A + (255 / 60)
		if a > 255 {
			a = 255
		}
	} else {
		a = i.CursorColor.A - (255 / 60)
		if a < 0 {
			a = 0
		}
	}

	i.CursorColor = color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: a}

	// handle cursor position
	width, _ := text.Measure(i.Data, i.Font, 0)

	i.CursorX = i.X + i.CursorOffsetX + width
	i.CursorY = i.Y + i.CursorOffsetY
}

func (i *Input) Draw(screen *ebiten.Image) {
	
}
