package bench

import (
	"bytes"
	"embed"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
	"image/draw"
	_ "image/png"
	"log"
)

//go:embed gopher.png
//go:embed gopher_greyscale.png
var staticFS embed.FS

func LoadSprite() *ebiten.Image {
	b, err := staticFS.ReadFile("gopher_greyscale.png")
	if err != nil {
		log.Fatal("Failed to open gopher.png file with err: ", err)
	}

	m, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return ebiten.NewImageFromImage(Checkerboard(25, 32, 4))
	}

	return ebiten.NewImageFromImage(m)
}

func Checkerboard(w, h, cells int) image.Image {
	m := image.NewRGBA(Rect(0, 0, w, h))
	cellW, cellH := w/cells, h/cells
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			c := color.RGBA{R: 255, B: 254, A: 255}
			if (i+j)%2 == 0 {
				c = color.RGBA{A: 255}
			}
			draw.Draw(m, Rect(i*cellW, j*cellH, cellW, cellH), &image.Uniform{C: c}, image.Point{}, draw.Src)
		}
	}
	return m
}

func Rect(x, y, w, h int) image.Rectangle {
	return image.Rect(x, y, x+w, y+h)
}
