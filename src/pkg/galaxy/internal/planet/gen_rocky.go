package planet

import (
	"math/rand"

	"github.com/mreliasen/ihniwiad/pkg/galaxy/public/consts"
	"github.com/mreliasen/ihniwiad/pkg/galaxy/public/interfaces"
)

/*
	Distribution before snow line
	-----------------------------------------------
	- 0.075 - 0.5 earth radii planets (15 % chance)
	- 0.5 - 1.5   earth radii planets (35% chance)
	- 1.5 - 2.5   earth radii planets (50 % chance)

	 rough guesstimates, extrapolated from our own solar system
*/

const (
	rocky_dwarf_min    float64 = 0.075 // earth radii
	rocky_dwarf_max    float64 = 0.49  // earth radii
	rocky_dwarf_chance float64 = 0.15  // % distribution of rocky planets within the frost line
)

const (
	rocky_earth_min    float64 = 0.5
	rocky_earth_max    float64 = 1.49
	rocky_earth_chance float64 = 0.35
)

const (
	rocky_super_min float64 = 1.5
	rocky_super_max float64 = 2.5
)

func GenerateRockyPlanet(rng *rand.Rand) interfaces.PlanetInterface {
	c := rng.Float64()
	m := rocky_super_min
	mx := rocky_super_max

	if c <= rocky_dwarf_chance {
		m = rocky_dwarf_min
		mx = rocky_dwarf_max
	} else if c <= rocky_earth_chance+rocky_dwarf_chance {
		m = rocky_earth_min
		mx = rocky_earth_max
	}

	return New(consts.ROCKY_PLANET, rng.Float64()*(mx-m)+m, 0.0)
}
