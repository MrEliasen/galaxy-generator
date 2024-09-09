package main

import (
	"github.com/charmbracelet/log"
	"github.com/mreliasen/ihniwiad/pkg/galaxy"
	"github.com/mreliasen/ihniwiad/pkg/logger"
	"github.com/mreliasen/ihniwiad/shared/utils"
)

func main() {
	logger := logger.New(log.DebugLevel)
	rng := utils.NewSeededRNG(12381)
	g := galaxy.New(rng, 50_000)

	i, o := g.HabitableZone()
	logger.Debugf("Galaxy Habitable Zone: %.0f - %.0f", i, o)

	g.GenerateStellarNeighbourhood()
	// g.Display()
}
