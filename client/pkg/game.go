package pkg

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jessehorne/endlessskyonline/client/pkg/ui"
)

type Game struct {
	UI           *ui.UI
	WindowWidth  int
	WindowHeight int
}

func NewGame() *Game {
	return &Game{
		UI:           ui.NewUI(),
		WindowWidth:  800,
		WindowHeight: 600,
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.UI.Draw(screen)
}

func (g *Game) Layout(w, h int) (int, int) {
	return g.WindowWidth, g.WindowHeight
}
