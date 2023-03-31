package pkg

import (
	"log"
	r "my-wolf3d/pkg/renderer"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Core struct {
	renderer *r.Renderer
}

func NewCore() *Core {
	return &Core{
		renderer: nil,
	}
}

func (c *Core) Init() {
	c.renderer = r.NewRenderer(screenWidth, screenHeight)
	c.renderer.Init()
}

func (c *Core) Run() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("My Wolf3d")

	//if err := ebiten.RunGame(&game.Game{Px: screenWidth, Py: screenHeight}); err != nil {
	if err := ebiten.RunGame(c.renderer); err != nil {
		log.Fatal(err)
	}
}

func (c *Core) Stop() {
	c.renderer = nil
}
