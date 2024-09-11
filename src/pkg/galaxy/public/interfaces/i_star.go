package interfaces

import "github.com/mreliasen/ihniwiad/pkg/coordinate/public/interfaces"

type StarInterface interface {
	/*
	   Returns the coordinates of the star (inherently the star system) within the galactic neighbourhood
	*/
	GetCoordinate() interfaces.CoordinateInterface

	/*
	   Returns the colour of the star
	*/
	GetColour() string

	/*
	   Returns the stars temperature in Kelvin
	*/
	GetTemperatureK() float64

	/*
	   Returns the star luminosity class/configuration
	*/
	GetLuminosityClass() string

	/*
	   Returns the stars sequence
	*/
	GetSequence() string

	/*
	   Returns the stars star classification
	*/
	GetClass() string

	/*
	   Returns the stars designation/name
	*/
	GetName() string

	/*
	   Returns the stars size in solar radii
	*/
	GetSize() float64

	/*
	   Returns the stars mass in solar masses
	*/
	GetMass() float64

	/*
	   Returns the stars mass in kg
	*/
	GetMassKg() float64

	/*
	   Returns the stars volume in m^3
	*/
	Volume() float64

	/*
	   Returns the stars Density in g/cm^3
	*/
	Density() float64

	/*
	   Returns the stars surface area in m^2
	*/
	Area() float64

	/*
	   Returns the stars luminosity in solar luminosities
	*/
	Luminosity() float64

	/*
	   Returns the stars Escape Velocity in m/s
	*/
	EscapeVelocity() (m float64)

	/*
	   Returns the stars Escape Velocity in m/s
	*/
	SolarWindVelocity() float64

	/*
	   Returns the stars solar wind mass loss rate in kg/s
	*/
	SolarWindMassLossRate() float64

	/*
	   Returns the stars thermal/termination shock distance in AU
	*/
	TerminalShockDistance() float64

	/*
	   Returns the stars frost line, in AU from the star
	*/
	FrostLine() float64

	/*
	   Returns habitable zone as start and end distance in AU from the star
	*/
	HabitableZone() (start, end float64)
}
