package planet

import (
	"math/rand"

	"github.com/mreliasen/ihniwiad/pkg/galaxy/public/consts"
	"github.com/mreliasen/ihniwiad/pkg/galaxy/public/interfaces"
)

// rough guesstimates, extrapolated from our own solar system
const (
	ice_size_min float64 = 2.5 // earth radii
	ice_size_max float64 = 4.0 // earth radii

	ice_density_min float64 = 0.8 // g/cm^3
	ice_density_max float64 = 2.4 // g/cm^3
)

func GenerateIceGiant(rng *rand.Rand) interfaces.PlanetInterface {
	return New(
		consts.ICE_GIANT,
		rng.Float64()*(ice_size_max-ice_size_min)+ice_size_min,
		rng.Float64()*(ice_density_max-ice_density_min)+ice_density_min,
	)
}
