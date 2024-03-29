package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"

	"github.com/chrisvaughn/hyrule-invaders/archetype"
	"github.com/chrisvaughn/hyrule-invaders/component"
)

type Controls struct {
	query *query.Query
}

func NewControls() *Controls {
	return &Controls{
		query: query.NewQuery(
			filter.Contains(
				transform.Transform,
				component.Input,
				component.Velocity,
				component.Sprite,
			),
		),
	}
}

func (i *Controls) Update(w donburi.World) {
	i.query.Each(w, func(entry *donburi.Entry) {
		input := component.Input.Get(entry)

		if input.Disabled {
			return
		}

		velocity := component.Velocity.Get(entry)

		velocity.Velocity = math.Vec2{
			X: 0,
			Y: 0,
		}

		if ebiten.IsKeyPressed(input.MoveUpKey) {
			velocity.Velocity.Y = -input.MoveSpeed
		} else if ebiten.IsKeyPressed(input.MoveDownKey) {
			velocity.Velocity.Y = input.MoveSpeed
		}

		if ebiten.IsKeyPressed(input.MoveRightKey) {
			velocity.Velocity.X = input.MoveSpeed
		}
		if ebiten.IsKeyPressed(input.MoveLeftKey) {
			velocity.Velocity.X = -input.MoveSpeed
		}

		player := archetype.MustFindPlayer(w)
		player.ShootTimer.Update()
		if inpututil.IsKeyJustPressed(input.ShootKey) && player.ShootTimer.IsReady() {
			position := transform.Transform.Get(entry).LocalPosition
			archetype.NewPlayerProjectile(w, position)
			player.ShootTimer.Reset()
		}
	})
}
