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
	for x := 0; x < g.vw; x++ {
		g.raycast(screen, x)
	}
}

func (g *Game) raycast(screen *ebiten.Image, x int) {
	rayAngle := g.camera.Angle() - g.camera.HalfFov()
	for rayCount := 0; rayCount < g.vw; rayCount++ {
		//mapX, mapY := g.camera.GetPosition()
		//
		//// Ray path incrementers
		//rayCos := math.Cos(utils.DegreeToRadians(rayAngle)) / float64(g.precision)
		//raySin := math.Sin(utils.DegreeToRadians(rayAngle)) / float64(g.precision)
		//
		//// Wall finder
		//for g.worldMap.HasHitWall(math.Floor(mapX), math.Floor(mapY)) {
		//	mapX += rayCos
		//	mapY += raySin
		//}
		//
		//// Pythagoras theorem
		//distance := math.Sqrt(math.Pow(g.camera.PosX()-mapX, 2) + math.Pow(g.camera.PosY()-mapY, 2))
		//
		//// Wall height
		//wallHeight := math.Floor(g.halfVHeight / distance)
		//
		////drawLine(rayCount, 0, rayCount, data.screen.halfHeight - wallHeight, "cyan");
		////drawLine(rayCount, data.screen.halfHeight - wallHeight, rayCount, data.screen.halfHeight + wallHeight, "red");
		////drawLine(rayCount, data.screen.halfHeight + wallHeight, rayCount, data.screen.height, "green");
		//
		//ebitenutil.DrawLine(screen, float64(rayCount), 0, float64(rayCount), g.halfVHeight, color.RGBA{R: 255, G: 0, B: 0, A: 100})
		//ebitenutil.DrawLine(screen, float64(rayCount), g.halfVHeight-wallHeight, float64(rayCount), g.halfVHeight+wallHeight, color.RGBA{R: 255, G: 255, B: 0, A: 100})
		//ebitenutil.DrawLine(screen, float64(rayCount), g.halfVHeight+wallHeight, float64(rayCount), float64(g.vh), color.RGBA{R: 255, G: 165, B: 0, A: 100})
		//
		//rayAngle += g.incrementAngle
	}
}

//func (g *Game) raycast(x int) {
//
//	posX, posY := g.camera.GetPosition()
//	dirX, dirY := g.camera.GetDirection()
//	planeX, planeY := g.camera.GetPlane()
//
//	//calculate ray position and direction
//	cameraX := float64(2*x/g.vw - 1) //x-coordinate in camera space
//	rayDirX := dirX + planeX*cameraX
//	rayDirY := dirY + planeY*cameraX
//
//	mapX := int(posX)
//	mapY := int(posY)
//
//	var sideDistX, sideDistY float64
//
//	deltaDistX := math.Abs(1 / rayDirX)
//	deltaDistY := math.Abs(1 / rayDirY)
//
//	var stepX, stepY int
//
//	var hit int = 0
//	var side int
//
//	//calculate step and initial sideDist
//	if rayDirX < 0 {
//		stepX = -1
//		sideDistX = (posX - float64(mapX)) * deltaDistX
//	} else {
//		stepX = 1
//		sideDistX = (float64(mapX) + 1.0 - posX) * deltaDistX
//	}
//
//	if rayDirY < 0 {
//		stepY = -1
//		sideDistY = (posY - float64(mapY)) * deltaDistY
//	} else {
//		stepY = 1
//		sideDistY = (float64(mapY) + 1.0 - posY) * deltaDistY
//	}
//
//	//perform DDA
//	for hit == 0 {
//		//jump to next map square, either in x-direction, or in y-direction
//		if sideDistX < sideDistY {
//			sideDistX += deltaDistX
//			mapX += stepX
//			side = 0
//		} else {
//			sideDistY += deltaDistY
//			mapY += stepY
//			side = 1
//		}
//		//Check if ray has hit a wall
//		if worldMap[mapX][mapY] > 0 {
//			hit = 1
//		}
//	}
//
//	//Calculate distance projected on camera direction (Euclidean distance would give fisheye effect!)
//	var perpWallDist float64
//	if side == 0 {
//		perpWallDist = sideDistX - deltaDistX
//	} else {
//		perpWallDist = sideDistY - deltaDistY
//	}
//
//	//Calculate height of line to draw on screen
//	var lineHeight int = int(float64(g.vh) / perpWallDist)
//
//	//calculate lowest and highest pixel to fill in current stripe
//	var drawStart int = -lineHeight/2 + g.vh/2
//	if drawStart < 0 {
//		drawStart = 0
//	}
//
//	var drawEnd int = lineHeight/2 + g.vh/2
//	if drawEnd >= g.vh {
//		drawEnd = g.vh - 1
//	}
//
//}
