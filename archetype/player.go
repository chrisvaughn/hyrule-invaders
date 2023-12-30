package archetype

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"

	"github.com/chrisvaughn/hyrule-invaders/assets"
	"github.com/chrisvaughn/hyrule-invaders/component"
	"github.com/chrisvaughn/hyrule-invaders/engine"
)

type PlayerInputs struct {
	Up    ebiten.Key
	Right ebiten.Key
	Down  ebiten.Key
	Left  ebiten.Key
	Shoot ebiten.Key
}

var inputs = PlayerInputs{
	Up:    ebiten.KeyUp,
	Right: ebiten.KeyRight,
	Down:  ebiten.KeyDown,
	Left:  ebiten.KeyLeft,
	Shoot: ebiten.KeySpace,
}

func playerSpawn(w donburi.World) math.Vec2 {
	game := component.MustFindGame(w)
	return math.Vec2{
		X: float64(game.Settings.ScreenWidth) * 0.50,
		Y: float64(game.Settings.ScreenHeight) * 0.9,
	}
}

func NewPlayer(w donburi.World) *donburi.Entry {

	player := component.PlayerData{
		Lives:        3,
		RespawnTimer: engine.NewTimer(time.Second * 3),
	}

	player.ShootTimer = engine.NewTimer(player.WeaponCooldown())

	return NewPlayerFromPlayerData(w, player)
}

func NewPlayerCharacter(w donburi.World) {
	character := w.Entry(
		w.Create(
			component.PlayerCharacter,
			transform.Transform,
			component.Velocity,
			component.Sprite,
			component.Input,
			component.Bounds,
			component.Collider,
		),
	)

	originalRotation := -90.0

	pos := playerSpawn(w)
	t := transform.Transform.Get(character)
	t.LocalPosition = pos
	t.LocalRotation = originalRotation

	component.Sprite.SetValue(character, component.SpriteData{
		Image:            assets.Link,
		Layer:            component.SpriteLayerCharacters,
		Pivot:            component.SpritePivotCenter,
		OriginalRotation: originalRotation,
	})

	s := assets.Link.Bounds().Size()
	component.Collider.SetValue(character, component.ColliderData{
		Width:  float64(s.X),
		Height: float64(s.Y),
		Layer:  component.CollisionLayerPlayers,
	})

	component.Input.SetValue(character, component.InputData{
		MoveUpKey:    inputs.Up,
		MoveRightKey: inputs.Right,
		MoveDownKey:  inputs.Down,
		MoveLeftKey:  inputs.Left,
		MoveSpeed:    3.5,
		ShootKey:     inputs.Shoot,
	})

}

func NewPlayerFromPlayerData(w donburi.World, playerData component.PlayerData) *donburi.Entry {
	player := w.Entry(w.Create(component.Player))
	component.Player.SetValue(player, playerData)
	return player
}

func MustFindPlayer(w donburi.World) *component.PlayerData {
	player, ok := query.NewQuery(filter.Contains(component.Player)).First(w)
	if !ok {
		panic("game not found")
	}
	return component.Player.Get(player)
}
