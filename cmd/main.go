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

type CoreGame struct {
	game *game.Game
}

func (g *CoreGame) Update() error {
	return nil
}

func (g *CoreGame) Draw(screen *ebiten.Image) {
	g.game.Render(screen)
}

func (g *CoreGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("My Wolf3d")

	game := game.NewGame(screenWidth, screenHeight)
	core := &CoreGame{
		game: game,
	}
	//if err := ebiten.RunGame(&game.Game{Px: screenWidth, Py: screenHeight}); err != nil {
	if err := ebiten.RunGame(core); err != nil {
		log.Fatal(err)
	}
}
