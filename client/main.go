package main

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/jessehorne/endlessskyonline/client/pkg"
	"github.com/jessehorne/endlessskyonline/client/pkg/ui/drawables"
	"log"
	"os"
)

func main() {
	game := pkg.NewGame()

	fileContent, err := os.ReadFile("./assets/Roboto-Italic-VariableFont_wdth,wght.ttf")
	if err != nil {
		log.Fatal(err)
	}

	fontSource, err := text.NewGoTextFaceSource(bytes.NewReader(fileContent))
	if err != nil {
		log.Fatal(err)
	}
	f := &text.GoTextFace{
		Source: fontSource,
		Size:   32,
	}

	t := drawables.NewText(100, 100, "Hello, World!")
	t.Font = f
	game.UI.GetActiveWindow().Add(t)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
