package archetype

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"

	"github.com/chrisvaughn/hyrule-invaders/assets"
	"github.com/chrisvaughn/hyrule-invaders/component"
)

func NewEnemy(
	w donburi.World,
	position math.Vec2,
	rotation float64,
) {
	airplane := w.Entry(
		w.Create(
			transform.Transform,
			component.Velocity,
			component.Sprite,
			component.Despawnable,
			component.Collider,
			component.Health,
		),
	)

	originalRotation := -90.0

	t := transform.Transform.Get(airplane)
	t.LocalPosition = position
	t.LocalRotation = originalRotation + rotation

	image := assets.OctorockLvl1
	component.Sprite.SetValue(airplane, component.SpriteData{
		Image:            image,
		Layer:            component.SpriteLayerCharacters,
		Pivot:            component.SpritePivotCenter,
		OriginalRotation: originalRotation,
	})

	s := image.Bounds().Size()

	component.Collider.SetValue(airplane, component.ColliderData{
		Width:  float64(s.X),
		Height: float64(s.Y),
		Layer:  component.CollisionLayerAirEnemies,
	})

	health := component.Health.Get(airplane)
	health.Health = 3
}
