package internal

import "github.com/mreliasen/ihniwiad/pkg/coordinate/public/interfaces"

func New(x, y, z float64) interfaces.CoordinateInterface {
	return &Coordinate{
		X: x,
		Y: y,
		Z: z,
	}
}

type Coordinate struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

func (c *Coordinate) GetZ() float64 {
	return c.Z
}

func (c *Coordinate) GetY() float64 {
	return c.Y
}

func (c *Coordinate) GetX() float64 {
	return c.X
}

func (c *Coordinate) SetZ(z float64) interfaces.CoordinateInterface {
	c.Z = z
	return c
}

func (c *Coordinate) SetY(y float64) interfaces.CoordinateInterface {
	c.Y = y
	return c
}

func (c *Coordinate) SetX(x float64) interfaces.CoordinateInterface {
	c.X = x
	return c
}

func (c *Coordinate) List() []float64 {
	return []float64{
		c.X,
		c.Y,
	}
}
