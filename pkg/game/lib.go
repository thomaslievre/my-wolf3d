package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	c "my-wolf3d/pkg/camera"
	w "my-wolf3d/pkg/map"
)

type Game struct {
	camera   *c.Camera
	worldMap *w.WorldMap

	//	viewport
	vw, vh                  int
	halfVWidth, halfVHeight float64

	// raycaster
	incrementAngle float64
	precision      int

	// render
	delay int
}

func NewGame(vw int, vh int) *Game {
	// Init worldmap
	worldmap := w.NewWorldMap()

	// Init camera
	camera := c.NewCamera()
	camera.SetPosX(12)
	camera.SetPosY(12)
	camera.SetDirX(-1)
	camera.SetDirY(0)
	camera.SetPlaneX(0)
	camera.SetPlaneY(0.66)

	return &Game{
		camera:         camera,
		worldMap:       worldmap,
		vh:             vh,
		vw:             vw,
		halfVWidth:     float64(vw / 2),
		halfVHeight:    float64(vh / 2),
		incrementAngle: camera.Fov() / float64(vw),
		precision:      64,
		delay:          30,
	}
}

func (g *Game) Render(screen *ebiten.Image) {
	for rayCount := 0; rayCount < g.vw; rayCount++ {

	}
}
