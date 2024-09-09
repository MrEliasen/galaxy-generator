package galaxy

import (
	"math/rand"

	"github.com/mreliasen/ihniwiad/pkg/galaxy/internal"
	"github.com/mreliasen/ihniwiad/pkg/galaxy/public/interfaces"
)

func New(rng *rand.Rand, rad float64) interfaces.GalaxyInterface {
	g := internal.Galaxy{}
	g.SetRadius(rad)
	g.SetRNG(rng)
	return &g
}
