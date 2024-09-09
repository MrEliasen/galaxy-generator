package coordinate

import (
	"github.com/mreliasen/ihniwiad/pkg/coordinate/internal"
	"github.com/mreliasen/ihniwiad/pkg/coordinate/public/interfaces"
)

func New(x, y float64) interfaces.CoordinateInterface {
	return internal.New(x, y, 0)
}

func New3D(x, y, z float64) interfaces.CoordinateInterface {
	return internal.New(x, y, z)
}
