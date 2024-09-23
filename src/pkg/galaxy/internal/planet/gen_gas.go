package planet

import (
	"math/rand"

	"github.com/mreliasen/ihniwiad/pkg/galaxy/public/consts"
	"github.com/mreliasen/ihniwiad/pkg/galaxy/public/interfaces"
)

// rough guesstimates, extrapolated from our own solar system
const (
	gas_size_min float64 = 5.0  // earth radii
	gas_size_max float64 = 14.0 // earth radii

	gas_density_min float64 = 0.2 // g/cm^3
	gas_density_max float64 = 2.2 // g/cm^3
)

func GenerateGasGiant(rng *rand.Rand) interfaces.PlanetInterface {
	return New(
		consts.GAS_GIANT,
		rng.Float64()*(gas_size_max-gas_size_min)+gas_size_min,
		rng.Float64()*(gas_density_max-gas_density_min)+gas_density_min,
	)
}
