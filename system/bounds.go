package system

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"

	"github.com/chrisvaughn/hyrule-invaders/component"
	"github.com/chrisvaughn/hyrule-invaders/engine"
)

type Bounds struct {
	query *query.Query
	game  *component.GameData
}

func NewBounds() *Bounds {
	return &Bounds{
		query: query.NewQuery(filter.Contains(
			component.PlayerCharacter,
			transform.Transform,
			component.Sprite,
			component.Bounds,
		)),
	}
}

func (b *Bounds) Update(w donburi.World) {
	if b.game == nil {
		b.game = component.MustFindGame(w)
		if b.game == nil {
			return
		}
	}

	b.query.Each(w, func(entry *donburi.Entry) {
		bounds := component.Bounds.Get(entry)
		if bounds.Disabled {
			return
		}

		t := transform.Transform.Get(entry)
		sprite := component.Sprite.Get(entry)

		s := sprite.Image.Bounds().Size()
		width, height := float64(s.Y), float64(s.X)

		var minX, maxX, minY, maxY float64

		switch sprite.Pivot {
		case component.SpritePivotTopLeft:
			minX = 0
			maxX = float64(b.game.Settings.ScreenWidth) - width

			minY = float64(b.game.Settings.ScreenHeight/3) * 2
			maxY = float64(b.game.Settings.ScreenHeight) - height
		case component.SpritePivotCenter:
			minX = width / 2
			maxX = float64(b.game.Settings.ScreenWidth) - width/2

			minY = float64(b.game.Settings.ScreenHeight/3) * 2
			maxY = float64(b.game.Settings.ScreenHeight) - height/2
		}

		t.LocalPosition.X = engine.Clamp(t.LocalPosition.X, minX, maxX)
		t.LocalPosition.Y = engine.Clamp(t.LocalPosition.Y, minY, maxY)
	})
}
