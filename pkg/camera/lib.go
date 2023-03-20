package camera

type Camera struct {
	posX, posY     float64
	dirX, dirY     float64
	planeX, planeY float64
}

func (c *Camera) PosX() float64 {
	return c.posX
}

func (c *Camera) SetPosX(posX float64) {
	c.posX = posX
}

func (c *Camera) PosY() float64 {
	return c.posY
}

func (c *Camera) SetPosY(posY float64) {
	c.posY = posY
}

func (c *Camera) DirX() float64 {
	return c.dirX
}

func (c *Camera) SetDirX(dirX float64) {
	c.dirX = dirX
}

func (c *Camera) DirY() float64 {
	return c.dirY
}

func (c *Camera) SetDirY(dirY float64) {
	c.dirY = dirY
}

func (c *Camera) PlaneX() float64 {
	return c.planeX
}

func (c *Camera) SetPlaneX(planeX float64) {
	c.planeX = planeX
}

func (c *Camera) PlaneY() float64 {
	return c.planeY
}

func (c *Camera) SetPlaneY(planeY float64) {
	c.planeY = planeY
}

func NewCamera() *Camera {
	return &Camera{}
}

func (c *Camera) GetPosition() (float64, float64) {
	return c.posX, c.posY
}

func (c *Camera) GetDirection() (float64, float64) {
	return c.dirX, c.dirY
}

func (c *Camera) GetPlane() (float64, float64) {
	return c.planeX, c.planeY
}
