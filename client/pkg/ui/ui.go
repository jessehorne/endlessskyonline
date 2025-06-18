package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type UI struct {
	Windows      map[string]*Window
	ActiveWindow string
}

func NewUI() *UI {
	return &UI{
		Windows: map[string]*Window{
			"default": NewWindow(),
		},
		ActiveWindow: "default",
	}
}

func (u *UI) Draw(screen *ebiten.Image) {
	u.GetActiveWindow().Draw(screen)
}

func (u *UI) GetActiveWindow() *Window {
	return u.Windows[u.ActiveWindow]
}
