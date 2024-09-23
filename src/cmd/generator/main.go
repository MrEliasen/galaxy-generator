package main

import (
	"github.com/charmbracelet/log"
	"github.com/mreliasen/ihniwiad/pkg/galaxy"
	"github.com/mreliasen/ihniwiad/pkg/logger"
	"github.com/mreliasen/ihniwiad/pkg/utils"
)

func main() {
	logger := logger.New(log.DebugLevel)
	var seed int64 = 12381
	rng := utils.NewSeededRNG(seed)
	g := galaxy.New(rng, seed)

	i, o := g.HabitableZone()
	logger.Debugf("Galaxy Habitable Zone: %.0f - %.0f", i, o)

	g.GenerateStellarNeighbourhood(rng.Int63())
	g.DumpStarSystems()
}
