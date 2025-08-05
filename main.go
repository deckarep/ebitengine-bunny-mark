package main

import (
	"github.com/sedyh/ebiten-bunny-mark/bench"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(800, 600)

	ebiten.SetVsyncEnabled(false)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(bench.NewGame(200, false)); err != nil {
		log.Fatal(err)
	}
}
