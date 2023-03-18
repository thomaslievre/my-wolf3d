package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"my-wolf3d/pkg/game"
)

const (
	screenWidth  = 640
	screenHeight = 480
	padding      = 20
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("My Wolf3d")
	if err := ebiten.RunGame(&game.Game{Px: screenWidth, Py: screenHeight}); err != nil {
		log.Fatal(err)
	}
}
