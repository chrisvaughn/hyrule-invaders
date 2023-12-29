package scene

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/chrisvaughn/hyrule-invaders/assets"
	"github.com/chrisvaughn/hyrule-invaders/engine"
)

type Title struct {
	screenWidth     int
	screenHeight    int
	newGameCallback func()
}

func NewTitle(screenWidth int, screenHeight int, newGameCallback func()) *Title {
	return &Title{
		screenWidth:     screenWidth,
		screenHeight:    screenHeight,
		newGameCallback: newGameCallback,
	}
}

func (t *Title) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) || ebiten.IsKeyPressed(ebiten.KeySpace) {
		t.newGameCallback()
		return
	}
}

func (t *Title) Draw(screen *ebiten.Image) {
	msg := "Hyrule Invaders"
	w, _ := engine.RenderedStringLength(msg, assets.NarrowFont)
	text.Draw(screen, msg, assets.NarrowFont, t.screenWidth/2-w/2, 250, color.White)
	msg = "Press space to start"
	w, _ = engine.RenderedStringLength(msg, assets.NarrowFont)
	text.Draw(screen, msg, assets.NarrowFont, t.screenWidth/2-w/2, 300, color.White)
}
