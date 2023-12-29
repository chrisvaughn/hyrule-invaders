package scene

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/yohamta/donburi"

	"github.com/chrisvaughn/hyrule-invaders/assets"
	"github.com/chrisvaughn/hyrule-invaders/component"
	"github.com/chrisvaughn/hyrule-invaders/engine"
)

type System interface {
	Update(w donburi.World)
}

type Drawable interface {
	Draw(w donburi.World, screen *ebiten.Image)
}

type Game struct {
	world     donburi.World
	systems   []System
	drawables []Drawable

	screenWidth  int
	screenHeight int
}

func NewGame(screenWidth int, screenHeight int) *Game {
	g := &Game{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
	}
	g.world = g.createWorld()

	return g
}

func (g *Game) createWorld() donburi.World {
	world := donburi.NewWorld()

	game := world.Entry(world.Create(component.Game))
	component.Game.SetValue(game, component.GameData{
		Score: 0,
		Settings: component.Settings{
			ScreenWidth:  g.screenWidth,
			ScreenHeight: g.screenHeight,
		},
	})

	return world
}

func (g *Game) Update() {
	gameData := component.MustFindGame(g.world)
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		gameData.Paused = !gameData.Paused
	}

	if gameData.Paused {
		return
	}

	for _, s := range g.systems {
		s.Update(g.world)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()

	gameData := component.MustFindGame(g.world)
	if gameData.Paused {
		msg := "Paused"
		w, _ := engine.RenderedStringLength(msg, assets.NarrowFont)
		text.Draw(screen, msg, assets.NarrowFont, g.screenWidth/2-w/2, 250, color.White)
	}

	for _, s := range g.drawables {
		s.Draw(g.world, screen)
	}
}
