package interfaces

type CoordinateInterface interface {
	GetY() float64
	GetX() float64
	GetZ() float64
	SetY(y float64) CoordinateInterface
	SetX(x float64) CoordinateInterface
	SetZ(z float64) CoordinateInterface
	List() []float64
}
