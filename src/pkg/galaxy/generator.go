package galaxy

import (
	"math/rand"

	"github.com/mreliasen/ihniwiad/pkg/galaxy/internal"
	"github.com/mreliasen/ihniwiad/pkg/galaxy/public/interfaces"
	"github.com/mreliasen/ihniwiad/pkg/utils"
)

func New(rng *rand.Rand, seed int64) interfaces.GalaxyInterface {
	radiiMin := 35_000.0
	radiiMax := 150_000.0

	thicknessMin := 800.0
	thicknessMax := 2200.0

	g := internal.Galaxy{
		Type: "Spiral",
	}
	g.SetRadius(utils.RoundFloat(rng.Float64()*(radiiMax-radiiMin)+radiiMin, 0))
	g.SetThickness(utils.RoundFloat(rng.Float64()*(thicknessMax-thicknessMin)+thicknessMin, 0))
	g.SetRNG(rng)
	g.SetSeed(seed)
	return &g
}
