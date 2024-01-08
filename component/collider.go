package component

import (
	"github.com/yohamta/donburi"
)

const (
	CollisionLayerPlayerProjectiles ColliderLayer = iota
	CollisionLayerEnemyProjectiles
	CollisionLayerPlayer
	CollisionLayerEnemies
)

type ColliderLayer int

type ColliderData struct {
	Width  float64
	Height float64
	Layer  ColliderLayer
}

var Collider = donburi.NewComponentType[ColliderData]()
