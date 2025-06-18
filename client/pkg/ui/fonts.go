package ui

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"os"
)

var Fonts map[string]*text.GoTextFaceSource
var TextFaces map[string]*text.GoTextFace

func loadDefaultFont(name, path string) error {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	fontSource, err := text.NewGoTextFaceSource(bytes.NewReader(fileContent))
	if err != nil {
		return err
	}

	Fonts[name] = fontSource

	return nil
}

func loadDefaultTextFace(fontName, name string, size float64) {
	f := &text.GoTextFace{
		Source: Fonts[fontName],
		Size:   size,
	}

	TextFaces[name] = f
}

func LoadDefaultFonts() error {
	Fonts = map[string]*text.GoTextFaceSource{}
	TextFaces = map[string]*text.GoTextFace{}

	// Load Roboto
	err := loadDefaultFont("roboto", "./assets/Roboto-VariableFont_wdth,wght.ttf")
	if err != nil {
		return err
	}

	loadDefaultTextFace("roboto", "roboto-32", 32)

	return nil
}
