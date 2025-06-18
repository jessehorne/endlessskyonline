package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jessehorne/endlessskyonline/client/pkg"
	"github.com/jessehorne/endlessskyonline/client/pkg/ui"
	"github.com/jessehorne/endlessskyonline/client/pkg/widgets"
	"log"
)

func main() {
	game := pkg.NewGame()
	err := game.Init()
	if err != nil {
		log.Fatal(err)
	}

	t := widgets.NewText(100, 100, "Hello, World!")
	t.Font = ui.TextFaces["roboto-32"]
	game.UI.GetActiveWindow().Add(t)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
