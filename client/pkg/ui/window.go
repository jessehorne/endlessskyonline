package ui

import "github.com/hajimehoshi/ebiten/v2"

type Window struct {
	Width     int
	Height    int
	Drawables map[Drawable]struct{}
}

const (
	defaultWidth  = 800
	defaultHeight = 600
)

func NewWindow() *Window {
	return &Window{
		Width:     defaultWidth,
		Height:    defaultHeight,
		Drawables: make(map[Drawable]struct{}),
	}
}

func (w *Window) Draw(screen *ebiten.Image) {
	for d := range w.Drawables {
		d.Draw(screen)
	}
}

func (w *Window) Add(d Drawable) {
	w.Drawables[d] = struct{}{}
}

func (w *Window) Remove(d Drawable) {
	delete(w.Drawables, d)
}
