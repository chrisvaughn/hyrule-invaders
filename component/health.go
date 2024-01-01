package component

import (
	"github.com/yohamta/donburi"
)

type HealthData struct {
	Health      int
	JustDamaged bool
}

func (d *HealthData) Damage() {
	if d.Health <= 0 {
		return
	}

	d.Health--
	d.JustDamaged = true
}

func (d *HealthData) HideDamageIndicator() {
	d.JustDamaged = false
}

var Health = donburi.NewComponentType[HealthData](HealthData{})
