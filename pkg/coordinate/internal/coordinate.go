package internal

import "github.com/mreliasen/ihniwiad/pkg/coordinate/public/interfaces"

func New(x, y, z float64) interfaces.CoordinateInterface {
	return &Coordinate{
		x,
		y,
		z,
	}
}

type Coordinate struct {
	x float64
	y float64
	z float64
}

func (c *Coordinate) GetZ() float64 {
	return c.z
}

func (c *Coordinate) GetY() float64 {
	return c.y
}

func (c *Coordinate) GetX() float64 {
	return c.x
}

func (c *Coordinate) SetZ(z float64) interfaces.CoordinateInterface {
	c.z = z
	return c
}

func (c *Coordinate) SetY(y float64) interfaces.CoordinateInterface {
	c.y = y
	return c
}

func (c *Coordinate) SetX(x float64) interfaces.CoordinateInterface {
	c.x = x
	return c
}

func (c *Coordinate) List() []float64 {
	return []float64{
		c.x,
		c.y,
	}
}
