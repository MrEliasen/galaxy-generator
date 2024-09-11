package interfaces

import "github.com/mreliasen/ihniwiad/pkg/galaxy/public/consts"

type PlanetInterface interface {
	GetType() consts.PLANET_TYPE
	/*
	   in kg
	*/
	GetMass() float64
	/*
	   in m^3
	*/
	GetVolume() float64
	/*
	   In g/cm^3
	*/
	GetDensity() float64
	/*
	   In earth radii
	*/
	GetSize() float64
	/*
		Set if planet is habitable / within the habitable zone
	*/
	SetHabitable(b bool)
	/*
	   Planet distance to the star, in AU
	*/
	GetOrbitDistance() float64
	/*
		Set distance to the star, in AU
	*/
	SetOrbitDistance(distanceAu float64)
	/*
		If planet is habitable / within the habitable zone
	*/
	IsHabitable() bool
	HasMoon() bool
	GetMoons() []PlanetInterface
}
