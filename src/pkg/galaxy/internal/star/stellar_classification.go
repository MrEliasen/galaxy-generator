package star

import (
	"math/rand"

	"github.com/mreliasen/ihniwiad/pkg/utils"
)

type StarClassType struct {
	Name         string
	TempRangeMin float64
	TempRangeMax float64
	maxPlanets   int
	gasGiantMin  float64
	gasGiantMax  float64
	iceGiantMin  float64
	iceGiantMax  float64
	rockyMin     float64
	rockyMax     float64
}

func (sc *StarClassType) RanRockyPlanetCount(rng *rand.Rand) int {
	return int(rng.Float64()*(sc.rockyMax-sc.rockyMin) + sc.rockyMin)
}

func (sc *StarClassType) RanGasGiantCount(rng *rand.Rand) int {
	return int(rng.Float64()*(sc.gasGiantMax-sc.gasGiantMin) + sc.gasGiantMin)
}

func (sc *StarClassType) RanIceGiantCount(rng *rand.Rand) int {
	return int(rng.Float64()*(sc.iceGiantMax-sc.iceGiantMin) + sc.iceGiantMin)
}

func (sc *StarClassType) GenerateStar(rng *rand.Rand) (*Star, error) {
	cl := calculateSpectralClass(rng, sc.Name)
	radius := cl.RadiusMin + rng.Float64()*(cl.RadiusMax-cl.RadiusMin)
	temp := sc.TempRangeMin + rng.Float64()*(sc.TempRangeMax-sc.TempRangeMin)

	if temp == 0 {
		temp = cl.TempRangeMin + rng.Float64()*(cl.TempRangeMax-cl.TempRangeMin)
	}

	radSteps := (cl.RadiusMax - cl.RadiusMin) / 100
	mass := (rng.Float64() * radSteps) * 100

	star := &Star{
		Class:        sc.Name,
		TemperatureK: utils.RoundFloat(temp, 0),
		SolarRadii:   utils.RoundFloat(radius, 5),
		SolarMasses:  utils.RoundFloat(mass, 2),
	}

	star.LuminosityClass = cl.Name
	star.Sequence = cl.Type
	star.Colour = cl.Color

	return star, nil
}

// The planet # is based on observational data and a whole bunch of theory.
// Is it anything but scientifically accurate
var (
	OClass = StarClassType{
		Name:         "O",
		TempRangeMin: 30000,
		TempRangeMax: 52000,
		maxPlanets:   6,
		gasGiantMin:  1.0,
		gasGiantMax:  3.0,
		iceGiantMin:  1.0,
		iceGiantMax:  3.0,
		rockyMin:     0.0,
		rockyMax:     2.0,
	}
	BClass = StarClassType{
		Name:         "B",
		TempRangeMin: 10000,
		TempRangeMax: 30000,
		maxPlanets:   7,
		gasGiantMin:  1.0,
		gasGiantMax:  3.0,
		iceGiantMin:  1.0,
		iceGiantMax:  3.0,
		rockyMin:     0.0,
		rockyMax:     2.0,
	}
	AClass = StarClassType{
		Name:         "A",
		TempRangeMin: 7500,
		TempRangeMax: 10000,
		maxPlanets:   9,
		gasGiantMin:  3.0,
		gasGiantMax:  6.0,
		iceGiantMin:  1.0,
		iceGiantMax:  3.0,
		rockyMin:     1.0,
		rockyMax:     3.0,
	}
	FClass = StarClassType{
		Name:         "F",
		TempRangeMin: 6000,
		TempRangeMax: 7500,
		maxPlanets:   12,
		gasGiantMin:  2.0,
		gasGiantMax:  5.0,
		iceGiantMin:  2.0,
		iceGiantMax:  4.0,
		rockyMin:     2.0,
		rockyMax:     4.0,
	}
	GClass = StarClassType{
		Name:         "G",
		TempRangeMin: 5200,
		TempRangeMax: 6000,
		maxPlanets:   10,
		gasGiantMin:  2.0,
		gasGiantMax:  3.0,
		iceGiantMin:  2.0,
		iceGiantMax:  3.0,
		rockyMin:     2.0,
		rockyMax:     5.0,
	}
	KClass = StarClassType{
		Name:         "K",
		TempRangeMin: 3700,
		TempRangeMax: 5200,
		maxPlanets:   9,
		gasGiantMin:  1.0,
		gasGiantMax:  2.0,
		iceGiantMin:  1.0,
		iceGiantMax:  3.0,
		rockyMin:     3.0,
		rockyMax:     5.0,
	}
	MClass = StarClassType{
		Name:         "M",
		TempRangeMin: 2400,
		TempRangeMax: 3700,
		maxPlanets:   8,
		gasGiantMin:  0.0,
		gasGiantMax:  1.0,
		iceGiantMin:  1.0,
		iceGiantMax:  2.0,
		rockyMin:     3.0,
		rockyMax:     6.0,
	}
	BDClass = StarClassType{
		Name:         "Brown Dwarf",
		TempRangeMin: 0, // we get this from the sub type
		TempRangeMax: 0, // we get this from the sub type
		maxPlanets:   2,
		gasGiantMin:  0.0,
		gasGiantMax:  1.0,
		iceGiantMin:  0.0,
		iceGiantMax:  1.0,
		rockyMin:     0.0,
		rockyMax:     1.0,
	}
	WDClass = StarClassType{
		Name:         "White Dwarf",
		TempRangeMin: 10_000,
		TempRangeMax: 100_000,
		maxPlanets:   3,
		gasGiantMin:  0.0,
		gasGiantMax:  1.0,
		iceGiantMin:  0.0,
		iceGiantMax:  1.0,
		rockyMin:     0.0,
		rockyMax:     2.0,
	}
)
