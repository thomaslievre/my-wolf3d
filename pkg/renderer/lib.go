package renderer

import (
	"fmt"
	"image/color"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
	//padding      = 20

	TILE_SIZE             = 64
	WALL_HEIGHT           = 64
	PROJECTIONPLANEWIDTH  = 640
	PROJECTIONPLANEHEIGHT = 480
	ANGLE60               = PROJECTIONPLANEWIDTH
	ANGLE30               = ANGLE60 / 2
	ANGLE15               = ANGLE30 / 2
	ANGLE90               = ANGLE30 * 3
	ANGLE180              = ANGLE90 * 2
	ANGLE270              = ANGLE90 * 3
	ANGLE360              = ANGLE60 * 6
	ANGLE0                = 0
	ANGLE5                = ANGLE30 / 6
	ANGLE10               = ANGLE5 * 2

	// player's attributes
	fPlayerX                            = 100
	fPlayerY                            = 160
	fPlayerArc                          = ANGLE0
	fPlayerDistanceToTheProjectionPlane = 277
	fPlayerHeight                       = 32
	fPlayerSpeed                        = 16
	fProjectionPlaneYCenter             = PROJECTIONPLANEHEIGHT / 2

	// map
	MAP_WIDTH  = 12
	MAP_HEIGHT = 12

	W byte = 1
	O byte = 0
)

var (
	// trigonometric tables
	fSinTable   [ANGLE360 + 1]float64
	fISinTable  [ANGLE360 + 1]float64
	fCosTable   [ANGLE360 + 1]float64
	fICosTable  [ANGLE360 + 1]float64
	fTanTable   [ANGLE360 + 1]float64
	fITanTable  [ANGLE360 + 1]float64
	fFishTable  [ANGLE360 + 1]float64
	fXStepTable [ANGLE360 + 1]float64
	fYStepTable [ANGLE360 + 1]float64

	// the following variables are used to keep the player coordinate in the overhead map
	fPlayerMapX, fPlayerMapY, fMinimapWidth int

	// map
	fMap = [MAP_HEIGHT][MAP_WIDTH]byte{
		{W, W, W, W, W, W, W, W, W, W, W, W},
		{W, O, O, O, O, O, O, O, O, O, O, W},
		{W, O, O, O, O, O, O, O, O, O, O, W},
		{W, O, O, O, O, O, O, O, O, O, O, W},
		{W, O, O, O, O, O, O, O, O, O, O, W},
		{W, O, O, O, O, O, O, O, O, O, O, W},
		{W, O, O, O, O, O, O, O, O, O, O, W},
		{W, O, O, W, W, W, O, O, O, O, O, W},
		{W, O, O, O, O, O, O, O, O, O, O, W},
		{W, O, O, O, O, O, O, O, O, O, O, W},
		{W, O, O, O, O, O, O, O, O, O, O, W},
		{W, W, W, W, W, W, W, W, W, W, W, W},
	}
)

type Renderer struct {
	viewportWidth, viewportHeight int
}

func NewRenderer(vw int, vh int) *Renderer {
	return &Renderer{
		viewportWidth:  vw,
		viewportHeight: vh,
	}
}

func (r *Renderer) Init() {
	createTables()
}

func (r *Renderer) Update() error {
	return nil
}

func (r *Renderer) Draw(screen *ebiten.Image) {
	render(screen)
	os.Exit(0)
}

func (r *Renderer) Layout(outsideWidth, outsideHeight int) (int, int) {
	return r.viewportWidth, r.viewportHeight
}

func arcToRad(arcAngle float64) float64 {
	return (arcAngle * math.Pi) / float64(ANGLE180)
}

func createTables() {
	var radian float64
	fSinTable = [ANGLE360 + 1]float64{}
	fISinTable = [ANGLE360 + 1]float64{}
	fCosTable = [ANGLE360 + 1]float64{}
	fICosTable = [ANGLE360 + 1]float64{}
	fTanTable = [ANGLE360 + 1]float64{}
	fITanTable = [ANGLE360 + 1]float64{}
	fFishTable = [ANGLE360 + 1]float64{}
	fXStepTable = [ANGLE360 + 1]float64{}
	fYStepTable = [ANGLE360 + 1]float64{}

	for i := 0; i <= ANGLE360; i++ {
		radian = arcToRad(float64(i)) + 0.0001
		fSinTable[i] = math.Sin(radian)
		fISinTable[i] = 1.0 / (fSinTable[i])
		fCosTable[i] = math.Cos(radian)
		fICosTable[i] = 1.0 / (fCosTable[i])
		fTanTable[i] = math.Tan(radian)
		fITanTable[i] = 1.0 / fTanTable[i]

		// FACING LEFT
		if i >= ANGLE90 && i < ANGLE270 {
			fXStepTable[i] = float64(TILE_SIZE) / fTanTable[i]
			if fXStepTable[i] > 0 {
				fXStepTable[i] = -fXStepTable[i]
			} else { // FACING RIGHT
				fXStepTable[i] = float64(TILE_SIZE) / fTanTable[i]
				if fXStepTable[i] < 0 {
					fXStepTable[i] = -fXStepTable[i]
				}
			}
		}
		// FACING DOWN
		if i >= ANGLE0 && i < ANGLE180 {
			fYStepTable[i] = float64(TILE_SIZE) * fTanTable[i]
			if fYStepTable[i] < 0 {
				fYStepTable[i] = -fYStepTable[i]

			}
		} else { // FACING UP
			fYStepTable[i] = float64(TILE_SIZE) * fTanTable[i]
			if fYStepTable[i] > 0 {
				fYStepTable[i] = -fYStepTable[i]
			}
		}
	}

	for i := -ANGLE30; i <= ANGLE30; i++ {
		radian = arcToRad(float64(i))
		// we don't have negative angle, so make it start at 0
		// this will give range 0 to 320
		fFishTable[i+ANGLE30] = 1.0 / math.Cos(radian)
	}
}

func drawBackground(screen *ebiten.Image) {
	// sky
	c := 25
	var r int
	var renderedColor color.RGBA
	for r := 0; r < PROJECTIONPLANEHEIGHT/2; r += 10 {
		renderedColor = color.RGBA{R: 255, G: 0, B: 0, A: 100}
		ebitenutil.DrawRect(screen, 0, float64(r), PROJECTIONPLANEWIDTH, 10, renderedColor)
		// fOffscreenGraphics.setColor(new Color(c, 125, 225));
		// fOffscreenGraphics.fillRect(0, r, PROJECTIONPLANEWIDTH, 10);
		c += 20
	}
	// ground
	c = 22
	for ; r < PROJECTIONPLANEHEIGHT; r += 15 {
		renderedColor = color.RGBA{R: 0, G: 0, B: 255, A: 100}
		ebitenutil.DrawRect(screen, 0, float64(r), PROJECTIONPLANEWIDTH, 15, renderedColor)

		// fOffscreenGraphics.setColor(new Color(c, 20, 20));
		// fOffscreenGraphics.fillRect(0, r, PROJECTIONPLANEWIDTH, 15);
		c += 15
	}
}

func drawOverheadMap(screen *ebiten.Image) {
	var renderedColor color.RGBA
	fMinimapWidth := 5
	for v := 0; v < MAP_HEIGHT; v++ {
		for u := 0; u < MAP_WIDTH; u++ {
			if fMap[v][u] == W {
				renderedColor = color.RGBA{R: 85, G: 255, B: 255, A: 100}
				//   fOffscreenGraphics.setColor(Color.cyan);
			} else {
				renderedColor = color.RGBA{R: 0, G: 0, B: 0, A: 100}
				//   fOffscreenGraphics.setColor(Color.black);
			}
			ebitenutil.DrawRect(screen, float64(PROJECTIONPLANEWIDTH+u*fMinimapWidth), float64(v*fMinimapWidth), float64(fMinimapWidth), float64(fMinimapWidth), renderedColor)
			// fOffscreenGraphics.fillRect(PROJECTIONPLANEWIDTH+(u*fMinimapWidth),
			//             (v*fMinimapWidth), fMinimapWidth, fMinimapWidth);
		}
	}
	fPlayerMapX = PROJECTIONPLANEWIDTH + fPlayerX/TILE_SIZE*fMinimapWidth
	fPlayerMapY = fPlayerY / TILE_SIZE * fMinimapWidth
	// fPlayerMapX=PROJECTIONPLANEWIDTH+(int)(((float)fPlayerX/(float)TILE_SIZE) * fMinimapWidth);
	// fPlayerMapY=(int)(((float)fPlayerY/(float)TILE_SIZE) * fMinimapWidth);
}

func render(screen *ebiten.Image) {
	var castArc, castColumn int
	var xGrid, yGrid int
	var distToNextXGrid, distToNextYGrid int
	var xIntersection, yIntersection float64
	var distToXGridBeingHit, distToYGridBeingHit float64
	var distToNextXIntersection, distToNextYIntersection float64
	var xGridIndex, yGridIndex int

	// drawBackground(screen)
	// drawOverheadMap(screen)

	castArc = fPlayerArc
	castArc -= ANGLE30

	if castArc < 0 {
		castArc = ANGLE360 + castArc
	}

	for castColumn = 0; castColumn < PROJECTIONPLANEWIDTH; castColumn += 5 {
		if castArc > ANGLE0 && castArc < ANGLE180 {
			xGrid = (fPlayerY/TILE_SIZE)*TILE_SIZE + TILE_SIZE
			distToNextXGrid = TILE_SIZE

			xtemp := fITanTable[castArc] * float64(yGrid-fPlayerY)
			xIntersection = xtemp + float64(fPlayerX)

		} else {
			yGrid = (fPlayerY / TILE_SIZE) * TILE_SIZE
			distToNextXGrid = -TILE_SIZE

			xtemp := fITanTable[castArc] * float64(xGrid-fPlayerY)
			xIntersection = xtemp + fPlayerX

			xGrid--
		}

		if castArc == ANGLE0 || castArc == ANGLE180 {
			distToXGridBeingHit = math.MaxFloat64
		} else {
			distToNextXIntersection = fXStepTable[castArc]
			for {
				xGridIndex = int(xIntersection / TILE_SIZE)
				// in the picture, yGridIndex will be 1
				yGridIndex = (xGrid / TILE_SIZE)

				if xGridIndex >= MAP_WIDTH || yGridIndex >= MAP_HEIGHT || xGridIndex < 0 || yGridIndex < 0 {
					distToXGridBeingHit = math.MaxFloat64
					break
				} else if fMap[yGridIndex][xGridIndex] != O {
					distToXGridBeingHit = (xIntersection - fPlayerX) * fICosTable[castArc]
					break
				} else { // else, the ray is not blocked, extend to the next block
					xIntersection += distToNextXIntersection
					xGrid += distToNextXGrid
				}
			}
		}

		// FOLLOW X RAY
		if castArc < ANGLE90 || castArc > ANGLE270 {
			yGrid = TILE_SIZE + (fPlayerX/TILE_SIZE)*TILE_SIZE
			distToNextYGrid = TILE_SIZE

			ytemp := fTanTable[castArc] * (float64(yGrid) - fPlayerX)
			yIntersection = ytemp + fPlayerY
		} else { // RAY FACING LEFT
			yGrid = (fPlayerX / TILE_SIZE) * TILE_SIZE
			distToNextYGrid = -TILE_SIZE

			ytemp := fTanTable[castArc] * (float64(yGrid) - fPlayerX)
			yIntersection = ytemp + fPlayerY

			yGrid--
		}

		// LOOK FOR VERTICAL WALL
		if castArc == ANGLE90 || castArc == ANGLE270 {
			distToYGridBeingHit = math.MaxFloat64
		} else {
			distToNextYIntersection = fYStepTable[castArc]
			for {
				// compute current map position to inspect
				xGridIndex = (yGrid / TILE_SIZE)
				yGridIndex = int(yIntersection / TILE_SIZE)

				if xGridIndex >= MAP_WIDTH || yGridIndex >= MAP_HEIGHT || xGridIndex < 0 || yGridIndex < 0 {
					distToYGridBeingHit = math.MaxFloat64
					break
				} else if fMap[yGridIndex][xGridIndex] != O {
					distToYGridBeingHit = (yIntersection - fPlayerY) * fISinTable[castArc]
					break
				} else {
					yIntersection += distToNextYIntersection
					yGrid += distToNextYGrid
				}
			}
		}

		// DRAW THE WALL SLICE
		// var scaleFactor float64
		var dist float64
		var topOfWall float64    // used to compute the top and bottom of the sliver that
		var bottomOfWall float64 // will be the staring point of floor and ceiling
		var renderedColor color.RGBA

		if distToXGridBeingHit < distToYGridBeingHit {
			// the next function call (drawRayOnMap()) is not a part of raycating rendering part,
			// it just draws the ray on the overhead map to illustrate the raycasting process
			//   drawRayOnOverheadMap(xIntersection, horizontalGrid);
			dist = distToXGridBeingHit
			renderedColor = color.RGBA{R: 180, G: 180, B: 180, A: 100}
			//   fOffscreenGraphics.setColor(Color.gray);
		} else {
			// the next function call (drawRayOnMap()) is not a part of raycating rendering part,
			// it just draws the ray on the overhead map to illustrate the raycasting process
			// drawRayOnOverheadMap(verticalGrid, yIntersection)
			dist = distToYGridBeingHit
			renderedColor = color.RGBA{R: 120, G: 120, B: 120, A: 100}
			// fOffscreenGraphics.setColor(Color.darkGray)
		}

		dist /= fFishTable[castColumn]
		// projected_wall_height/wall_height = fPlayerDistToProjectionPlane/dist;
		projectedWallHeight := WALL_HEIGHT * fPlayerDistanceToTheProjectionPlane / dist
		bottomOfWall = float64(fProjectionPlaneYCenter) + projectedWallHeight*0.5
		topOfWall = PROJECTIONPLANEHEIGHT - bottomOfWall
		if bottomOfWall >= PROJECTIONPLANEHEIGHT {
			bottomOfWall = PROJECTIONPLANEHEIGHT - 1
		}

		//fOffscreenGraphics.drawLine(castColumn, topOfWall, castColumn, bottomOfWall);

		fmt.Println(castColumn, " | ", topOfWall, " | ", bottomOfWall, " | ", projectedWallHeight)
		ebitenutil.DrawLine(screen, float64(castColumn), float64(topOfWall), float64(castColumn), float64(bottomOfWall), renderedColor)
		// ebitenutil.DrawRect(screen, float64(castColumn), float64(topOfWall), 5, projectedWallHeight, renderedColor)
		//   fOffscreenGraphics.fillRect(castColumn, topOfWall, 5, projectedWallHeight);

		// TRACE THE NEXT RAY
		castArc += 5
		if castArc >= ANGLE360 {
			castArc -= ANGLE360
		}

	}

}
