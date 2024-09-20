package internal

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/mreliasen/ihniwiad/pkg/coordinate"
	ic "github.com/mreliasen/ihniwiad/pkg/coordinate/public/interfaces"
	"github.com/mreliasen/ihniwiad/pkg/galaxy/internal/star"
	"github.com/mreliasen/ihniwiad/pkg/galaxy/public/interfaces"
)

type StellarNeighbourhood struct {
	Rng     *rand.Rand                       `json:"-"`
	Seed    int64                            `json:"seed"`
	Dist    float64                          `json:"distance"`
	Density float64                          `json:"stellar_density"`
	Radius  float64                          `json:"neighbourhood_radius"`
	Coords  ic.CoordinateInterface           `json:"galactic_coordinate"`
	Systems []interfaces.StarSystemInterface `json:"star_systems"`
}

func (sn *StellarNeighbourhood) GetRadius() float64 {
	return sn.Radius
}

func (sn *StellarNeighbourhood) GetSystems() []interfaces.StarSystemInterface {
	return sn.Systems
}

func (sn *StellarNeighbourhood) PopulateNeighbourhood() {
	vol := (4.0 / 3.0) * math.Pi * math.Pow(sn.Radius, 3)
	starCount := int(vol * sn.Density)

	systems := []interfaces.StarSystemInterface{}

	dsn := (sn.Coords.GetX() + sn.Coords.GetY())
	if dsn < 0 {
		dsn *= -1
	}

	// Generate stellar mass object placeholders
	for i := 0; i < starCount; i++ {
		config := sn.randomStarClassification()
		star, err := config.GenerateStar(sn.Rng)
		if err != nil {
			continue
		}

		star.Name = fmt.Sprintf("%.0f.%d.%s", dsn, len(systems)+1, star.GetClass())

		star.Coordinate = coordinate.New3D(
			sn.Rng.Float64()*sn.Radius*2-sn.Radius,
			sn.Rng.Float64()*sn.Radius*2-sn.Radius,
			sn.Rng.Float64()*sn.Radius*2-sn.Radius,
		)

		system := &StarSystem{}
		system.SetRNG(sn.Rng)
		system.SetStar(star)
		system.SetRockyPlants(config.RanRockyPlanetCount(sn.Rng))
		system.SetIceGiants(config.RanIceGiantCount(sn.Rng))
		system.SetGasGiants(config.RanGasGiantCount(sn.Rng))
		system.PopulatePlanets()

		systems = append(systems, system)
	}

	sn.Systems = systems
}

func (sn *StellarNeighbourhood) randomStarClassification() star.StarClassType {
	num := sn.Rng.Float64() * 100

	switch true {
	case num <= 3.0e-005:
		return star.OClass
	case num <= 0.13:
		return star.BClass
	case num <= 0.6:
		return star.AClass
	case num <= 3.0:
		return star.FClass
	case num <= 7.6:
		return star.GClass
	case num <= 12.1:
		return star.KClass
	default:
		return star.MClass
	}
}
