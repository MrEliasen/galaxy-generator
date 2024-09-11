package planet

import (
	"math"

	"github.com/mreliasen/ihniwiad/pkg/galaxy/public/consts"
	"github.com/mreliasen/ihniwiad/pkg/galaxy/public/interfaces"
	"github.com/mreliasen/ihniwiad/pkg/utils"
)

type Planet struct {
	PlanetType consts.PLANET_TYPE           `json:"planet_type"`
	Habitable  bool                         `json:"habitable"`
	Moons      []interfaces.PlanetInterface `json:"moons,omitempty"`
	EarthRadii float64                      `json:"earth_radii"`
	Density    float64                      `json:"density"`
	OrbitAU    float64                      `json:"orbit_au"`
}

func (p *Planet) GetType() consts.PLANET_TYPE {
	return p.PlanetType
}

/*
in kg
*/
func (p *Planet) GetMass() float64 {
	if p.PlanetType != consts.ROCKY_PLANET {
		return (p.Density * 1000) * p.GetVolume()
	}

	return math.Pow(p.EarthRadii, 3.7) * consts.EARTH_MASS
}

/*
If planet is habitable / within the habitable zone
*/
func (p *Planet) IsHabitable() bool {
	return p.Habitable
}

/*
Set if planet is habitable / within the habitable zone
*/
func (p *Planet) SetHabitable(b bool) {
	p.Habitable = b
}

/*
Set distance to the star, in AU
*/
func (p *Planet) SetOrbitDistance(distanceAu float64) {
	p.OrbitAU = utils.RoundFloat(distanceAu, 3)
}

/*
Planet distance to the star, in AU
*/
func (p *Planet) GetOrbitDistance() float64 {
	return p.OrbitAU
}

/*
in m^3
*/
func (p *Planet) GetVolume() float64 {
	return (4.0 / 3.0) * math.Pi * math.Pow(p.EarthRadii*consts.EARTH_RADII, 3)
}

/*
In g/cm^3
*/
func (p *Planet) GetDensity() float64 {
	if p.PlanetType != consts.ROCKY_PLANET {
		return p.Density
	}

	return (p.GetMass() / p.GetVolume()) / 1000
}

/*
In earth radii
*/
func (p *Planet) GetSize() float64 {
	return p.EarthRadii
}

func (p *Planet) HasMoon() bool {
	return len(p.Moons) > 0
}

func (p *Planet) GetMoons() []interfaces.PlanetInterface {
	return p.Moons
}
