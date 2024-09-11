package planet

import (
	"github.com/mreliasen/ihniwiad/pkg/galaxy/public/consts"
	"github.com/mreliasen/ihniwiad/pkg/galaxy/public/interfaces"
	"github.com/mreliasen/ihniwiad/pkg/utils"
)

func New(planetType consts.PLANET_TYPE, earthRadii, density float64) interfaces.PlanetInterface {
	p := Planet{
		PlanetType: planetType,
		Habitable:  false,
		Moons:      []interfaces.PlanetInterface{},
		EarthRadii: utils.RoundFloat(earthRadii, 3),
		Density:    density,
	}

	p.Density = utils.RoundFloat(p.GetDensity(), 4)

	return &p
}
