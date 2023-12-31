package system

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"golang.org/x/image/colornames"

	"github.com/chrisvaughn/hyrule-invaders/component"
)

type Debug struct {
	query     *query.Query
	debug     *component.DebugData
	offscreen *ebiten.Image

	restartLevelCallback func()
}

func NewDebug(restartLevelCallback func()) *Debug {
	return &Debug{
		query: query.NewQuery(
			filter.Contains(transform.Transform, component.Sprite),
		),
		// TODO figure out the proper size
		offscreen:            ebiten.NewImage(3000, 3000),
		restartLevelCallback: restartLevelCallback,
	}
}

func (d *Debug) Update(w donburi.World) {
	if d.debug == nil {
		debug, ok := query.NewQuery(filter.Contains(component.Debug)).First(w)
		if !ok {
			return
		}

		d.debug = component.Debug.Get(debug)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySlash) {
		d.debug.Enabled = !d.debug.Enabled
	}

	if d.debug.Enabled {
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			d.restartLevelCallback()
		}
	}
}

func (d *Debug) Draw(w donburi.World, screen *ebiten.Image) {
	if d.debug == nil || !d.debug.Enabled {
		return
	}

	allCount := w.Len()

	despawnableCount := 0
	spawnedCount := 0
	query.NewQuery(filter.Contains(component.Despawnable)).Each(w, func(entry *donburi.Entry) {
		despawnableCount++
		if component.Despawnable.Get(entry).Spawned {
			spawnedCount++
		}
	})

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Entities: %v Despawnable: %v Spawned: %v", allCount, despawnableCount, spawnedCount), 0, 0)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %v", ebiten.ActualTPS()), 0, 20)

	d.offscreen.Clear()
	d.query.Each(w, func(entry *donburi.Entry) {
		t := transform.Transform.Get(entry)
		sprite := component.Sprite.Get(entry)

		position := transform.WorldPosition(entry)

		s := sprite.Image.Bounds().Size()
		halfW, halfH := float64(s.X)/2, float64(s.Y)/2

		x := position.X
		y := position.Y

		switch sprite.Pivot { // nolint:exhaustive
		case component.SpritePivotCenter:
			x -= halfW
			y -= halfH
		}

		vector.DrawFilledRect(d.offscreen, float32(t.LocalPosition.X-2), float32(t.LocalPosition.Y-2), 4, 4, colornames.Lime, false)
		ebitenutil.DebugPrintAt(d.offscreen, fmt.Sprintf("%v", entry.Entity().Id()), int(x), int(y))
		ebitenutil.DebugPrintAt(d.offscreen, fmt.Sprintf("pos: %.0f, %.0f", position.X, position.Y), int(x), int(y)+40)
		ebitenutil.DebugPrintAt(d.offscreen, fmt.Sprintf("rot: %.0f", transform.WorldRotation(entry)), int(x), int(y)+60)

		length := 50.0
		right := position.Add(transform.Right(entry).MulScalar(length))
		up := position.Add(transform.Up(entry).MulScalar(length))

		vector.StrokeLine(d.offscreen, float32(position.X), float32(position.Y), float32(right.X), float32(right.Y), 1, colornames.Blue, false)
		vector.StrokeLine(d.offscreen, float32(position.X), float32(position.Y), float32(up.X), float32(up.Y), 1, colornames.Lime, false)

		if entry.HasComponent(component.Collider) {
			collider := component.Collider.Get(entry)
			vector.StrokeLine(d.offscreen, float32(x), float32(y), float32(x+collider.Width), float32(y), 1, colornames.Lime, false)
			vector.StrokeLine(d.offscreen, float32(x), float32(y), float32(x), float32(y+collider.Height), 1, colornames.Lime, false)
			vector.StrokeLine(d.offscreen, float32(x+collider.Width), float32(y), float32(x+collider.Width), float32(y+collider.Height), 1, colornames.Lime, false)
			vector.StrokeLine(d.offscreen, float32(x), float32(y+collider.Height), float32(x+collider.Width), float32(y+collider.Height), 1, colornames.Lime, false)
		}
	})

	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(d.offscreen, op)
}
