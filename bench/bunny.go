package bench

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	gravity = 0.00095
)

type Bunny struct {
	Hue        float32
	PosX, PosY float32
	VelX, VelY float32
}

func NewBunny(shift, hue float32) *Bunny {
	return &Bunny{
		Hue:  hue,
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

func (b *Bunny) Draw(screen *ebiten.Image, sprite *ebiten.Image, colorful bool) {
	sw, sh := float32(screen.Bounds().Dx()), float32(screen.Bounds().Dy())

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.PosX*sw), float64(b.PosY*sh))
	if colorful {
		op.ColorM.RotateHue(float64(b.Hue))
	}
	screen.DrawImage(sprite, op)
}
