package archetype

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"

	"github.com/chrisvaughn/hyrule-invaders/assets"
	"github.com/chrisvaughn/hyrule-invaders/component"
)

const (
	playerProjectileSpeed = 10
)

func NewPlayerProjectile(w donburi.World, position math.Vec2) {
	width := float64(assets.Arrow.Bounds().Dy())

	newPlayerProjectile(w, math.Vec2{
		X: position.X,
		Y: position.Y - width,
	}, 0)

}

func newPlayerProjectile(w donburi.World, position math.Vec2, localRotation float64) {
	projectile := w.Entry(
		w.Create(
			component.Velocity,
			transform.Transform,
			component.Sprite,
			component.Despawnable,
			component.Collider,
		),
	)

	image := assets.Arrow

	originalRotation := -90.0

	t := transform.Transform.Get(projectile)
	t.LocalPosition = position
	t.LocalRotation = originalRotation + localRotation

	component.Velocity.SetValue(projectile, component.VelocityData{
		Velocity: transform.Right(projectile).MulScalar(playerProjectileSpeed),
	})

	component.Sprite.SetValue(projectile, component.SpriteData{
		Image:            image,
		Layer:            component.SpriteLayerCharacters,
		Pivot:            component.SpritePivotCenter,
		OriginalRotation: originalRotation,
	})

	width, height := image.Size()
	component.Collider.SetValue(projectile, component.ColliderData{
		Width:  float64(width),
		Height: float64(height),
		Layer:  component.CollisionLayerPlayerBullets,
	})
}
