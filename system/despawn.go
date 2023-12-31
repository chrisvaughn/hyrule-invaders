package system

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/hierarchy"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"

	"github.com/chrisvaughn/hyrule-invaders/component"
)

type Despawn struct {
	query *query.Query
	game  *component.GameData
}

func NewDespawn() *Despawn {
	return &Despawn{
		query: query.NewQuery(filter.Contains(component.Despawnable)),
	}
}

func (d *Despawn) Update(w donburi.World) {
	if d.game == nil {
		d.game = component.MustFindGame(w)
		if d.game == nil {
			return
		}
	}

	d.query.Each(w, func(entry *donburi.Entry) {
		position := transform.Transform.Get(entry).LocalPosition
		sprite := component.Sprite.Get(entry)
		despawnable := component.Despawnable.Get(entry)

		maxX := position.X + float64(sprite.Image.Bounds().Dx())
		maxY := position.Y + float64(sprite.Image.Bounds().Dy())

		if !despawnable.Spawned {
			if position.Y > 0 && maxY < float64(d.game.Settings.ScreenHeight) &&
				position.X > 0 && maxX < float64(d.game.Settings.ScreenWidth) {
				despawnable.Spawned = true
			}

			return
		}

		if maxY < 0 || position.Y > float64(d.game.Settings.ScreenHeight) ||
			maxX < 0 || position.X > float64(d.game.Settings.ScreenWidth) {
			hierarchy.RemoveRecursive(entry)
		}
	})
}
