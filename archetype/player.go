package archetype

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"

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

type PlayerSettings struct {
	Inputs PlayerInputs
}

var Players = map[int]PlayerSettings{
	1: {
		Inputs: PlayerInputs{
			Up:    ebiten.KeyW,
			Right: ebiten.KeyD,
			Down:  ebiten.KeyS,
			Left:  ebiten.KeyA,
			Shoot: ebiten.KeySpace,
		},
	},
	2: {
		Inputs: PlayerInputs{
			Up:    ebiten.KeyUp,
			Right: ebiten.KeyRight,
			Down:  ebiten.KeyDown,
			Left:  ebiten.KeyLeft,
			Shoot: ebiten.KeyEnter,
		},
	},
}

func NewPlayer(w donburi.World) *donburi.Entry {

	player := component.PlayerData{
		Lives:        3,
		RespawnTimer: engine.NewTimer(time.Second * 3),
	}

	//player.ShootTimer = engine.NewTimer(player.WeaponCooldown())

	return NewPlayerFromPlayerData(w, player)
}

func NewPlayerFromPlayerData(w donburi.World, playerData component.PlayerData) *donburi.Entry {
	player := w.Entry(w.Create(component.Player))
	component.Player.SetValue(player, playerData)
	return player
}
