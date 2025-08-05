package bench

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
)

const (
	gravity = 0.00095
)

var (
	gopherColor = color.RGBA{R: 71, G: 255, B: 234, A: 255}
)

type Bunny struct {
	Hue        int32
	PosX, PosY float32
	VelX, VelY float32
}

func NewBunny(shift float32, hueIndex int32) Bunny {
	return Bunny{
		Hue:  hueIndex,
		PosX: shift,
		VelX: float32(RangeFloat(0, 0.005)),
		VelY: float32(RangeFloat(0.0025, 0.005)),
	}
}

func (b *Bunny) Update(sprite *ebiten.Image, bounds image.Rectangle) {
	b.PosX += b.VelX
	b.PosY += b.VelY
	b.VelY += gravity

	sw, sh := float32(bounds.Dx()), float32(bounds.Dy())
	iw, ih := float32(sprite.Bounds().Dx()), float32(sprite.Bounds().Dy())
	relW, relH := iw/sw, ih/sh

	if b.PosX+relW > 1 {
		b.VelX *= -1
		b.PosX = 1 - relW
	}
	if b.PosX < 0 {
		b.VelX *= -1
		b.PosX = 0
	}
	if b.PosY+relH > 1 {
		b.VelY *= -0.85
		b.PosY = 1 - relH
		if Chance(0.5) {
			b.VelY -= float32(RangeFloat(0, 0.009))
		}
	}
	if b.PosY < 0 {
		b.VelY = 0
		b.PosY = 0
	}
}

func (b *Bunny) Draw(screen *ebiten.Image, sprite *ebiten.Image, colorful bool, colorSelection []color.Color) {
	sb := screen.Bounds()
	sw, sh := float32(sb.Dx()), float32(sb.Dy())

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.PosX*sw), float64(b.PosY*sh))
	if colorful {
		op.ColorScale.ScaleWithColor(colorSelection[b.Hue])
		screen.DrawImage(sprite, op)
	} else {
		op.ColorScale.ScaleWithColor(gopherColor)
		screen.DrawImage(sprite, op)
	}
}
