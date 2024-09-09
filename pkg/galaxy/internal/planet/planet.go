package planet

import (
	"math"

	"github.com/mreliasen/ihniwiad/pkg/galaxy/public/consts"
	"github.com/mreliasen/ihniwiad/pkg/galaxy/public/interfaces"
)

type Planet struct {
	planetType consts.PLANET_TYPE
	habitable  bool
	moons      []interfaces.PlanetInterface
	earthRadii float64
	density    float64
	orbitAu    float64
}

func (p *Planet) GetType() consts.PLANET_TYPE {
	return p.planetType
}

/*
in kg
*/
func (p *Planet) GetMass() float64 {
	if p.planetType != consts.ROCKY_PLANET {
		return (p.density * 1000) * p.GetVolume()
	}

	return math.Pow(p.earthRadii, 3.7) * consts.EARTH_MASS
}

/*
If planet is habitable / within the habitable zone
*/
func (p *Planet) IsHabitable() bool {
	return p.habitable
}

/*
Set if planet is habitable / within the habitable zone
*/
func (p *Planet) SetHabitable(b bool) {
	p.habitable = b
}

/*
Set distance to the star, in AU
*/
func (p *Planet) SetOrbitDistance(distanceAu float64) {
	p.orbitAu = distanceAu
}

/*
Planet distance to the star, in AU
*/
func (p *Planet) GetOrbitDistance() float64 {
	return p.orbitAu
}

/*
in m^3
*/
func (p *Planet) GetVolume() float64 {
	return (4.0 / 3.0) * math.Pi * math.Pow(p.earthRadii*consts.EARTH_RADII, 3)
}

/*
In g/cm^3
*/
func (p *Planet) GetDensity() float64 {
	if p.planetType != consts.ROCKY_PLANET {
		return p.density
	}

	return (p.GetMass() / p.GetVolume()) / 1000
}

/*
In earth radii
*/
func (p *Planet) GetSize() float64 {
	return p.earthRadii
}

func (p *Planet) HasMoon() bool {
	return len(p.moons) > 0
}

func (p *Planet) GetMoons() []interfaces.PlanetInterface {
	return p.moons
}
