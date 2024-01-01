package assets

import (
	"bytes"
	"embed"
	_ "embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	//go:embed */*
	assetsFS embed.FS

	/* Fonts Section */
	SmallFont  font.Face
	NormalFont font.Face
	NarrowFont font.Face

	/* Characters Section */
	Link         *ebiten.Image
	OctorockLvl1 *ebiten.Image

	/* Objects Section */
	Arrow        *ebiten.Image
	Health       *ebiten.Image
	OctorockRock *ebiten.Image
)

func MustLoadAssets() {
	normalFontData := mustLoadFile(assetsFS, "fonts/kenney-future.ttf")
	narrowFontData := mustLoadFile(assetsFS, "fonts/kenney-future-narrow.ttf")
	SmallFont = mustLoadFont(normalFontData, 10)
	NormalFont = mustLoadFont(normalFontData, 24)
	NarrowFont = mustLoadFont(narrowFontData, 24)

	/* Characters Section */
	Link = mustNewEbitenImage(mustLoadFile(assetsFS, "characters/link.png"))
	OctorockLvl1 = mustNewEbitenImage(mustLoadFile(assetsFS, "characters/octorock_lvl_1.png"))

	/* Objects Section */
	Health = mustNewEbitenImage(mustLoadFile(assetsFS, "objects/heart.png"))
	Arrow = mustNewEbitenImage(mustLoadFile(assetsFS, "objects/arrow.png"))
	OctorockRock = mustNewEbitenImage(mustLoadFile(assetsFS, "objects/octorock_rock.png"))
}

func mustLoadFile(fs embed.FS, filename string) []byte {
	data, err := fs.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return data
}

func mustLoadFont(data []byte, size int) font.Face {
	f, err := opentype.Parse(data)
	if err != nil {
		panic(err)
	}

	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    float64(size),
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
	}

	return face
}

func mustNewEbitenImage(data []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
