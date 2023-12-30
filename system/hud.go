package system

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"golang.org/x/image/colornames"

	"github.com/chrisvaughn/hyrule-invaders/assets"
	"github.com/chrisvaughn/hyrule-invaders/component"
	"github.com/chrisvaughn/hyrule-invaders/engine"
)

type HUD struct {
	query         *query.Query
	game          *component.GameData
	shadowOverlay *ebiten.Image
}

func NewHUD() *HUD {
	return &HUD{
		query: query.NewQuery(filter.Contains(component.Player)),
	}
}

func (h *HUD) Draw(w donburi.World, screen *ebiten.Image) {
	if h.game == nil {
		h.game = component.MustFindGame(w)
		if h.game == nil {
			return
		}
		h.shadowOverlay = ebiten.NewImage(h.game.Settings.ScreenWidth, h.game.Settings.ScreenHeight)
		h.shadowOverlay.Fill(colornames.Black)
	}

	h.query.Each(w, func(entry *donburi.Entry) {
		player := component.Player.Get(entry)

		icon := assets.Health
		iconHeight := icon.Bounds().Dy()

		baseY := float64(h.game.Settings.ScreenHeight) - float64(iconHeight) - 5
		var baseX float64 = 5

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(baseX, baseY)
		for i := 0; i < player.Lives; i++ {
			if i > 0 {
				op.GeoM.Translate(0, -float64(iconHeight+2))
			}
			screen.DrawImage(icon, op)
		}
	})

	if h.game.GameOver {
		op := &ebiten.DrawImageOptions{}
		op.ColorScale.Scale(0, 0, 0, 0.5)
		screen.DrawImage(h.shadowOverlay, op)

		msg := "GAME OVER"
		strWidth, _ := engine.RenderedStringLength(msg, assets.NormalFont)
		text.Draw(
			screen,
			msg,
			assets.NormalFont,
			h.game.Settings.ScreenWidth/2-strWidth/2,
			h.game.Settings.ScreenHeight/2,
			colornames.White,
		)
	} else if h.game.Paused {
		op := &ebiten.DrawImageOptions{}
		op.ColorScale.Scale(0, 0, 0, 0.5)
		screen.DrawImage(h.shadowOverlay, op)

		msg := "PAUSED"
		strWidth, _ := engine.RenderedStringLength(msg, assets.NormalFont)
		text.Draw(
			screen,
			msg,
			assets.NormalFont,
			h.game.Settings.ScreenWidth/2-strWidth/2,
			h.game.Settings.ScreenHeight/2,
			colornames.White,
		)
	}
	msg := fmt.Sprintf("Score: %06d", h.game.Score)
	strWidth, _ := engine.RenderedStringLength(msg, assets.NormalFont)
	text.Draw(screen, msg, assets.NormalFont, h.game.Settings.ScreenWidth/2-strWidth/2, 30, colornames.White)
}
