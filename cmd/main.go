package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const (
	screenWidth  = 640
	screenHeight = 480
	padding      = 20
)

type CoreGame struct {
}

func (g *CoreGame) Update() error {
	return nil
}

func (g *CoreGame) Draw(screen *ebiten.Image) {
	fmt.Println(screen.Size())
}

func (g *CoreGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("My Wolf3d")
	//if err := ebiten.RunGame(&game.Game{Px: screenWidth, Py: screenHeight}); err != nil {
	if err := ebiten.RunGame(&CoreGame{}); err != nil {
		log.Fatal(err)
	}
}
