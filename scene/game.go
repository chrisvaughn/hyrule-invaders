package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type System interface {
	Update(w donburi.World)
}

type Drawable interface {
	Draw(w donburi.World, screen *ebiten.Image)
}

type Game struct{}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Update() {}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
}
