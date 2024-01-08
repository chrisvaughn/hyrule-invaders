package system

import (
	"fmt"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/hierarchy"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"

	"github.com/chrisvaughn/hyrule-invaders/archetype"
	"github.com/chrisvaughn/hyrule-invaders/component"
	"github.com/chrisvaughn/hyrule-invaders/engine"
)

type Collision struct {
	query *query.Query
}

func NewCollision() *Collision {
	return &Collision{
		query: query.NewQuery(filter.Contains(component.Collider)),
	}
}

type collisionEffect func(w donburi.World, entry *donburi.Entry, other *donburi.Entry)

var collisionEffects = map[component.ColliderLayer]map[component.ColliderLayer]collisionEffect{
	component.CollisionLayerPlayerProjectiles: {
		component.CollisionLayerEnemies: func(w donburi.World, entry *donburi.Entry, other *donburi.Entry) {
			hierarchy.RemoveRecursive(entry)
			component.Health.Get(other).Damage()
		},
	},
	component.CollisionLayerEnemies: {
		component.CollisionLayerPlayer: func(w donburi.World, entry *donburi.Entry, other *donburi.Entry) {
			EnemyKilledEvent.Publish(w, EnemyKilled{
				Enemy: entry,
			})
			damagePlayer(w, other)
		},
	},
	component.CollisionLayerEnemyProjectiles: {
		component.CollisionLayerPlayer: func(w donburi.World, entry *donburi.Entry, other *donburi.Entry) {
			hierarchy.RemoveRecursive(entry)
			damagePlayer(w, other)
		},
	},
}

func damagePlayer(w donburi.World, entry *donburi.Entry) {
	hierarchy.RemoveRecursive(entry)

	player := archetype.MustFindPlayer(w)
	player.Damage()
}

func (c *Collision) Update(w donburi.World) {
	var entries []*donburi.Entry
	c.query.Each(w, func(entry *donburi.Entry) {
		// Skip entities not spawned yet
		if entry.HasComponent(component.Despawnable) {
			if !component.Despawnable.Get(entry).Spawned {
				return
			}
		}
		entries = append(entries, entry)
	})

	for _, entry := range entries {
		if !entry.Valid() {
			continue
		}

		collider := component.Collider.Get(entry)

		for _, other := range entries {
			if entry.Entity().Id() == other.Entity().Id() {
				continue
			}

			// One of the entities could already be removed from the world due to collision effect
			if !entry.Valid() || !other.Valid() {
				continue
			}

			otherCollider := component.Collider.Get(other)

			effects, ok := collisionEffects[collider.Layer]
			if !ok {
				continue
			}

			effect, ok := effects[otherCollider.Layer]
			if !ok {
				continue
			}

			if !entry.HasComponent(transform.Transform) {
				panic(fmt.Sprintf("%#v missing position\n", entry.Entity().Id()))
			}
			pos := transform.Transform.Get(entry).LocalPosition
			otherPos := transform.Transform.Get(other).LocalPosition

			rect := engine.NewRect(pos.X, pos.Y, collider.Width, collider.Height)
			otherRect := engine.NewRect(otherPos.X, otherPos.Y, otherCollider.Width, otherCollider.Height)

			if rect.Intersects(otherRect) {
				effect(w, entry, other)
			}
		}
	}
}
