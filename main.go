package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/chrisvaughn/hyrule-invaders/assets"
	"github.com/chrisvaughn/hyrule-invaders/scene"
)

const (
	screenWidth  = 800
	screenHeight = 640
)

type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
}

type Game struct {
	scene Scene
}

func NewGame() *Game {
	assets.MustLoadAssets()

	g := &Game{}
	g.switchToTitle()
	return g
}

func (g *Game) switchToTitle() {
	g.scene = scene.NewTitle(screenWidth, screenHeight, g.switchToGame)
}

func (g *Game) switchToGame() {
	g.scene = scene.NewGame(screenWidth, screenHeight)
}

func (g *Game) Update() error {
	g.scene.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.scene.Draw(screen)
}

func (g *Game) Layout(width, height int) (int, int) {
	return width, height
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)

	err := ebiten.RunGame(NewGame())
	if err != nil {
		log.Fatal(err)
	}
}
