package planet

import (
	"github.com/mreliasen/ihniwiad/pkg/galaxy/public/consts"
	"github.com/mreliasen/ihniwiad/pkg/galaxy/public/interfaces"
)

func New(planetType consts.PLANET_TYPE, earthRadii, density float64) interfaces.PlanetInterface {
	return &Planet{
		planetType: planetType,
		habitable:  false,
		moons:      []interfaces.PlanetInterface{},
		earthRadii: earthRadii,
		density:    density,
	}
}
