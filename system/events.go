package system

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/events"
	"github.com/yohamta/donburi/features/hierarchy"

	"github.com/chrisvaughn/hyrule-invaders/component"
)

type EnemyKilled struct {
	Enemy *donburi.Entry
}

var EnemyKilledEvent = events.NewEventType[EnemyKilled]()

func OnEnemyKilledAddScore(w donburi.World, event EnemyKilled) {
	component.MustFindGame(w).AddScore(1)
}

func OnEnemyKilledDestroyEnemy(w donburi.World, event EnemyKilled) {
	hierarchy.RemoveRecursive(event.Enemy)
}

func SetupEvents(w donburi.World) {
	EnemyKilledEvent.Subscribe(w, OnEnemyKilledAddScore)
	EnemyKilledEvent.Subscribe(w, OnEnemyKilledDestroyEnemy)
}

type Events struct{}

func NewEvents() *Events {
	return &Events{}
}

func (e *Events) Update(w donburi.World) {
	EnemyKilledEvent.ProcessEvents(w)
}
