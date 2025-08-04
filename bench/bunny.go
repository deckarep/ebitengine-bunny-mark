package bench

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
	"image"
	"image/color"
)

const (
	gravity             = 0.00095
	origR, origG, origB = 71.0 / 255.0, 255.0 / 255.0, 234 / 255.0
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

func (b *Bunny) Draw(screen *ebiten.Image, sprite *ebiten.Image, colorful bool) {
	sb := screen.Bounds()
	sw, sh := float32(sb.Dx()), float32(sb.Dy())

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.PosX*sw), float64(b.PosY*sh))
	if colorful {
		var c color.Color
		switch b.Hue {
		case 0:
			c = colornames.Blue
		case 1:
			c = colornames.Green
		case 2:
			c = colornames.Purple
		case 3:
			c = colornames.Yellow
		default:

			c = colornames.Red
		}

		_r, _g, _b, _ := c.RGBA()
		op.ColorM.Scale(float64(_r), float64(_g), float64(_b), 1.0)
		screen.DrawImage(sprite, op)
	} else {
		op.ColorM.Scale(origR, origG, origB, 1.0)
		screen.DrawImage(sprite, op)
	}
}
