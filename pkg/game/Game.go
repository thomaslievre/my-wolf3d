package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	"math"
	"strings"
)

const (
	mapWidth  = 24
	mapHeight = 24
)

var worldMap = [mapWidth][mapHeight]int{
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 2, 2, 2, 2, 2, 0, 0, 0, 0, 3, 0, 3, 0, 3, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 2, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 2, 0, 0, 0, 2, 0, 0, 0, 0, 3, 0, 0, 0, 3, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 2, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 2, 2, 0, 2, 2, 0, 0, 0, 0, 3, 0, 3, 0, 3, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 4, 4, 4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 4, 0, 4, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 4, 0, 0, 0, 0, 5, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 4, 0, 4, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 4, 0, 4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 4, 4, 4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
}

var posX, posY float64 = 4, 4
var dirX, dirY float64 = 0, -1
var planeX, planeY float64 = 0, 0.66
var time float64 = 0
var oldTime float64 = 0

type Game struct {
	Px, Py int
	keys   []ebiten.Key
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for x := 0; x < g.Px; x++ {
		//calculate ray position and direction
		cameraX := float64(2*x/g.Px - 1) //x-coordinate in camera space
		rayDirX := dirX + planeX*cameraX
		rayDirY := dirY + planeY*cameraX

		mapX := int(posX)
		mapY := int(posY)

		var sideDistX, sideDistY float64

		deltaDistX := math.Abs(1 / rayDirX)
		deltaDistY := math.Abs(1 / rayDirY)

		var stepX, stepY int

		var hit int = 0
		var side int

		//calculate step and initial sideDist
		if rayDirX < 0 {
			stepX = -1
			sideDistX = (posX - float64(mapX)) * deltaDistX
		} else {
			stepX = 1
			sideDistX = (float64(mapX) + 1.0 - posX) * deltaDistX
		}

		if rayDirY < 0 {
			stepY = -1
			sideDistY = (posY - float64(mapY)) * deltaDistY
		} else {
			stepY = 1
			sideDistY = (float64(mapY) + 1.0 - posY) * deltaDistY
		}

		//perform DDA
		for hit == 0 {
			//jump to next map square, either in x-direction, or in y-direction
			if sideDistX < sideDistY {
				sideDistX += deltaDistX
				mapX += stepX
				side = 0
			} else {
				sideDistY += deltaDistY
				mapY += stepY
				side = 1
			}
			//Check if ray has hit a wall
			if worldMap[mapX][mapY] > 0 {
				hit = 1
			}
		}

		//Calculate distance projected on camera direction (Euclidean distance would give fisheye effect!)
		var perpWallDist float64
		if side == 0 {
			perpWallDist = sideDistX - deltaDistX
		} else {
			perpWallDist = sideDistY - deltaDistY
		}

		//Calculate height of line to draw on screen
		var lineHeight int = int(float64(g.Py) / perpWallDist)

		//calculate lowest and highest pixel to fill in current stripe
		var drawStart int = -lineHeight/2 + g.Py/2
		if drawStart < 0 {
			drawStart = 0
		}

		var drawEnd int = lineHeight/2 + g.Py/2
		if drawEnd >= g.Py {
			drawEnd = g.Py - 1
		}

		var renderedColor color.RGBA

		switch worldMap[mapX][mapY] {
		case 1:
			renderedColor = color.RGBA{R: 255, G: 0, B: 0, A: 100}
			break
		case 2:
			renderedColor = color.RGBA{R: 0, G: 255, B: 0, A: 100}
			break
		case 3:
			renderedColor = color.RGBA{R: 0, G: 0, B: 255, A: 100}
			break
		case 4:
			renderedColor = color.RGBA{R: 255, G: 255, B: 255, A: 100}
			break
		default:
			renderedColor = color.RGBA{R: 255, G: 165, B: 0, A: 100}
			break
		}

		if side == 1 {
			renderedColor.A = renderedColor.A / 2
		}

		ebitenutil.DrawLine(screen, float64(x), float64(drawStart), float64(x), float64(drawEnd), renderedColor)
	}

	oldTime = time
	time =

	//keyStrs := []string{}
	//for _, k := range g.keys {
	//	keyStrs = append(keyStrs, k.String())
	//}
	//ebitenutil.DebugPrint(screen, strings.Join(keyStrs, ", "))

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Px, g.Py
}
