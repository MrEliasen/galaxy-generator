package star

import (
	"math/rand"

	"github.com/mreliasen/ihniwiad/pkg/utils"
)

type StarClassType struct {
	Name         string
	SizeMin      float64
	SizeMax      float64
	MassMin      float64
	MassMax      float64
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
	temp := sc.TempRangeMin + rng.Float64()*(sc.TempRangeMax-sc.TempRangeMin)

	if temp < 2400 {
		temp = 2400
	}

	star := &Star{
		Class:        sc.Name,
		SolarRadii:   utils.RoundFloat(rng.Float64()*(sc.SizeMax-sc.SizeMin)+sc.SizeMin, 4),
		SolarMasses:  utils.RoundFloat(rng.Float64()*(sc.MassMax-sc.MassMin)+sc.MassMin, 4),
		TemperatureK: utils.RoundFloat(temp, 0),
	}

	cl, err := calculateSpectralClass(star)
	if err != nil {
		return nil, err
	}

	star.LuminosityClass = cl.Name
	star.Sequence = cl.Type
	star.Colour = cl.Color

	return star, nil
}

var (
	MClass = StarClassType{
		Name:         "M",
		SizeMin:      0.1,
		SizeMax:      0.7,
		MassMin:      0.1,
		MassMax:      0.6,
		TempRangeMin: 2400,
		TempRangeMax: 3699,
		maxPlanets:   8,
		gasGiantMin:  0.0,
		gasGiantMax:  1.0,
		iceGiantMin:  1.0,
		iceGiantMax:  2.0,
		rockyMin:     3.0,
		rockyMax:     6.0,
	}
	KClass = StarClassType{
		Name:         "K",
		SizeMin:      0.8,
		SizeMax:      0.9,
		MassMin:      0.6,
		MassMax:      0.8,
		TempRangeMin: 3700,
		TempRangeMax: 5199,
		maxPlanets:   9,
		gasGiantMin:  1.0,
		gasGiantMax:  2.0,
		iceGiantMin:  1.0,
		iceGiantMax:  3.0,
		rockyMin:     3.0,
		rockyMax:     5.0,
	}
	GClass = StarClassType{
		Name:         "G",
		SizeMin:      0.9,
		SizeMax:      1.1,
		MassMin:      0.8,
		MassMax:      1.2,
		TempRangeMin: 5200,
		TempRangeMax: 5999,
		maxPlanets:   10,
		gasGiantMin:  2.0,
		gasGiantMax:  3.0,
		iceGiantMin:  2.0,
		iceGiantMax:  3.0,
		rockyMin:     2.0,
		rockyMax:     5.0,
	}
	FClass = StarClassType{
		Name:         "F",
		SizeMin:      1.1,
		SizeMax:      1.4,
		MassMin:      1.2,
		MassMax:      1.4,
		TempRangeMin: 6000,
		TempRangeMax: 7499,
		maxPlanets:   12,
		gasGiantMin:  2.0,
		gasGiantMax:  5.0,
		iceGiantMin:  2.0,
		iceGiantMax:  4.0,
		rockyMin:     2.0,
		rockyMax:     4.0,
	}
	AClass = StarClassType{
		Name:         "A",
		SizeMin:      1.4,
		SizeMax:      1.8,
		MassMin:      1.4,
		MassMax:      2.0,
		TempRangeMin: 7500,
		TempRangeMax: 9999,
		maxPlanets:   9,
		gasGiantMin:  3.0,
		gasGiantMax:  6.0,
		iceGiantMin:  1.0,
		iceGiantMax:  3.0,
		rockyMin:     1.0,
		rockyMax:     3.0,
	}
	BClass = StarClassType{
		Name:         "B",
		SizeMin:      1.8,
		SizeMax:      6.6,
		MassMin:      2.0,
		MassMax:      15.0,
		TempRangeMin: 10000,
		TempRangeMax: 29999,
		maxPlanets:   7,
		gasGiantMin:  1.0,
		gasGiantMax:  3.0,
		iceGiantMin:  1.0,
		iceGiantMax:  3.0,
		rockyMin:     0.0,
		rockyMax:     2.0,
	}
	OClass = StarClassType{
		Name:         "O",
		SizeMin:      6.6,
		SizeMax:      1700.0,
		MassMin:      15.0,
		MassMax:      30.0,
		TempRangeMin: 30000,
		TempRangeMax: 60000,
		maxPlanets:   6,
		gasGiantMin:  1.0,
		gasGiantMax:  3.0,
		iceGiantMin:  1.0,
		iceGiantMax:  3.0,
		rockyMin:     0.0,
		rockyMax:     2.0,
	}
)
