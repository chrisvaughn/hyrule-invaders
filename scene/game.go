package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"

	"github.com/chrisvaughn/hyrule-invaders/archetype"
	"github.com/chrisvaughn/hyrule-invaders/component"
	"github.com/chrisvaughn/hyrule-invaders/system"
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

	g.loadLevel()

	return g
}

func (g *Game) loadLevel() {
	render := system.NewRenderer()
	debug := system.NewDebug(g.loadLevel)

	g.systems = []System{
		system.NewControls(),
		system.NewVelocity(),
		system.NewBounds(),
		system.NewDespawn(),
		render,
		debug,
	}

	g.drawables = []Drawable{
		render,
		debug,
		system.NewHUD(),
	}

	g.world = g.createWorld()
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

	archetype.NewPlayer(world)
	archetype.NewPlayerCharacter(world)

	world.Create(component.Debug)

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
	for _, s := range g.drawables {
		s.Draw(g.world, screen)
	}
}
