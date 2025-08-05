package main

import (
	"github.com/sedyh/ebiten-bunny-mark/bench"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(bench.NewGame(200, false)); err != nil {
		log.Fatal(err)
	}
}
