package bench

import (
	"image"
	"image/color"
	"math/rand"

	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
)

type Game struct {
	Sprite         *ebiten.Image    // Image for bunnies
	Bounds         *image.Rectangle // Physical window size
	Bunnies        []Bunny          // List of bunnies
	Amount         *int             // How much to add
	Metrics        *Metrics         // Current TPS, FPS, object count and plots
	Colorful       *bool            // Add some serious load
	ColorSelection []color.Color
	Gpu            string // Current gpu
}

func NewGame(amount int, colorful bool) *Game {
	g := &Game{
		Sprite:         LoadSprite(),
		Amount:         &amount,
		Colorful:       &colorful,
		ColorSelection: make([]color.Color, 0),
		Bounds:         &image.Rectangle{},
		Bunnies:        make([]Bunny, 0, 200_000),
	}

	g.Metrics = NewMetrics(500*time.Millisecond, g.Bounds, g.Colorful, g.Amount)

	// Setup some colors.
	g.ColorSelection = append(g.ColorSelection, color.RGBA{R: 243, G: 174, B: 250, A: 255})
	g.ColorSelection = append(g.ColorSelection, color.RGBA{R: 184, G: 253, B: 150, A: 255})
	g.ColorSelection = append(g.ColorSelection, color.RGBA{R: 240, G: 226, B: 131, A: 255})
	g.ColorSelection = append(g.ColorSelection, color.RGBA{R: 245, G: 194, B: 165, A: 255})
	g.ColorSelection = append(g.ColorSelection, color.RGBA{R: 209, G: 200, B: 251, A: 255})

	g.AddBunnies()

	return g
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.AddBunnies()
	}
	//if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
	//	g.AddBunnies()
	//}

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
		g.Bunnies[i].Draw(screen, g.Sprite, *g.Colorful, g.ColorSelection)
	}

	g.Metrics.Draw(screen)
}

func (g *Game) Layout(width, height int) (int, int) {
	g.Bounds.Max = image.Point{X: width, Y: height}

	return width, height
}

func (g *Game) AddBunnies() {
	numColors := len(g.ColorSelection)

	for i := 0; i < *g.Amount; i++ {
		b := NewBunny(
			float32(len(g.Bunnies)%2),
			int32(rand.Intn(numColors)),
		)
		g.Bunnies = append(g.Bunnies, b)
	}
}

func (g *Game) RemoveBunnies() {
	if len(g.Bunnies) > 0 {
		g.Bunnies = g.Bunnies[0 : len(g.Bunnies)-*g.Amount]
	}
}
