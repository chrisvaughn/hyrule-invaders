package component

import (
	"time"

	"github.com/yohamta/donburi"

	"github.com/chrisvaughn/hyrule-invaders/engine"
)

type PlayerData struct {
	Lives        int
	Respawning   bool
	RespawnTimer *engine.Timer
	ShootTimer   *engine.Timer
}

func (d *PlayerData) AddLive() {
	d.Lives++
}

func (d *PlayerData) Damage() {
	if d.Respawning {
		return
	}

	d.Lives--

	if d.Lives > 0 {
		d.Respawning = true
		d.RespawnTimer.Reset()
	}
}

func (d *PlayerData) WeaponCooldown() time.Duration {
	return 400 * time.Millisecond
}

var Player = donburi.NewComponentType[PlayerData]()
var PlayerCharacter = donburi.NewComponentType[struct{}]()
