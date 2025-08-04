package bench

import (
	"image"
	"math/rand"

	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
)

type Game struct {
	Sprite   *ebiten.Image    // Image for bunnies
	Bounds   *image.Rectangle // Physical window size
	Bunnies  []Bunny          // List of bunnies
	Amount   *int             // How much to add
	Metrics  *Metrics         // Current TPS, FPS, object count and plots
	Colorful *bool            // Add some serious load
	Gpu      string           // Current gpu
}

func NewGame(amount int, colorful bool) *Game {
	g := &Game{
		Sprite:   LoadSprite(),
		Amount:   &amount,
		Colorful: &colorful,
		Bounds:   &image.Rectangle{},
		Bunnies:  make([]Bunny, 0, 100_000),
	}

	g.Metrics = NewMetrics(500*time.Millisecond, g.Bounds, g.Colorful, g.Amount)
	g.AddBunnies()

	return g
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.AddBunnies()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDelete) {
		g.RemoveBunnies()
	}

	if ids := ebiten.AppendTouchIDs(nil); len(ids) > 0 {
		g.AddBunnies() // not accurate, cause no input manager for this
	}

	if _, offset := ebiten.Wheel(); offset != 0 {
		*g.Amount += int(offset * 10)
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		*g.Colorful = !*g.Colorful
	}

	for i := 0; i < len(g.Bunnies); i++ {
		g.Bunnies[i].Update(g.Sprite, *g.Bounds)
	}

	g.Metrics.Update(float64(len(g.Bunnies)))

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Whitesmoke)

	for i := 0; i < len(g.Bunnies); i++ {
		g.Bunnies[i].Draw(screen, g.Sprite, *g.Colorful)
	}

	g.Metrics.Draw(screen)
}

func (g *Game) Layout(width, height int) (int, int) {
	g.Bounds.Max = image.Point{X: width, Y: height}

	return width, height
}

func (g *Game) AddBunnies() {
	for i := 0; i < *g.Amount; i++ {
		b := NewBunny(
			float32(len(g.Bunnies)%2),
			int32(rand.Intn(6)),
		)

		g.Bunnies = append(g.Bunnies, b)
	}
}

func (g *Game) RemoveBunnies() {
	if len(g.Bunnies) > 0 {
		g.Bunnies = g.Bunnies[0 : len(g.Bunnies)-*g.Amount]
	}
}
