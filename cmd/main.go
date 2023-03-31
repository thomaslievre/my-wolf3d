package main

import (
	pkg "my-wolf3d/pkg"
)

func main() {
	// ebiten.SetWindowSize(screenWidth, screenHeight)
	// ebiten.SetWindowTitle("My Wolf3d")

	// createTables()
	// //game := game.NewGame(screenWidth, screenHeight)
	// //	game: game,
	// core := &CoreGame{}
	// //if err := ebiten.RunGame(&game.Game{Px: screenWidth, Py: screenHeight}); err != nil {
	// if err := ebiten.RunGame(core); err != nil {
	// 	log.Fatal(err)
	// }

	core := pkg.NewCore()
	core.Init()
	core.Run()
	core.Stop()
}
