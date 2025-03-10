package assets

import (
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

// Embed the assets directory
//
//go:embed kenney_simple-space/PNG/Default/*
var content embed.FS

// Load player sprite
var PlayerSprite = mustLoadImage("kenney_simple-space/PNG/Default/ship_E.png")
var PlayerEngineEffect = mustLoadImage("kenney_simple-space/PNG/Default/effect_purple.png")

func mustLoadImage(name string) *ebiten.Image {
	f, err := content.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
