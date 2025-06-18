package ui

import "github.com/hajimehoshi/ebiten/v2"

type Drawable interface {
	Draw(*ebiten.Image)
}
