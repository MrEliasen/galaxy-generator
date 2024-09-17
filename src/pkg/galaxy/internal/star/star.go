package star

import (
	"encoding/json"
	"math"
	"math/rand"

	ic "github.com/mreliasen/ihniwiad/pkg/coordinate/public/interfaces"
	"github.com/mreliasen/ihniwiad/pkg/galaxy/public/consts"
	"github.com/mreliasen/ihniwiad/pkg/utils"
)

type Star struct {
	Rng             *rand.Rand             `json:"-"`
	Name            string                 `json:"designation"`
	Coordinate      ic.CoordinateInterface `json:"neighbourhood_coordinates"`
	Class           string                 `json:"class"`
	LuminosityClass string                 `json:"luminosity_class"`
	Colour          string                 `json:"colour"`
	Sequence        string                 `json:"sequence"`
	SolarRadii      float64                `json:"solar_radii"`
	SolarMasses     float64                `json:"solar_masses"`
	TemperatureK    float64                `json:"temperature_k"`
}

func (s Star) MarshalJSON() ([]byte, error) {
	// Create a type alias to avoid infinite recursion
	type Alias Star

	hzI, hzO := s.HabitableZone()

	return json.Marshal(&struct {
		Alias
		FrostLine     float64   `json:"frost_line"`
		HabitableZone []float64 `json:"habitable_zone"`
	}{
		Alias:         (Alias)(s),
		FrostLine:     utils.RoundFloat(s.FrostLine(), 2),
		HabitableZone: []float64{utils.RoundFloat(hzI, 2), utils.RoundFloat(hzO, 2)},
	})
}

/*
Returns the coordinates of the star (inherently the star system) within the galactic neighbourhood
*/
func (star *Star) GetCoordinate() ic.CoordinateInterface {
	return star.Coordinate
}

/*
Returns the colour of the star
*/
func (star *Star) GetColour() string {
	return star.Colour
}

/*
Returns the stars temperature in Kelvin
*/
func (star *Star) GetTemperatureK() float64 {
	return star.TemperatureK
}

/*
Returns the star luminosity class/configuration
*/
func (star *Star) GetLuminosityClass() string {
	return star.LuminosityClass
}

/*
Returns the stars sequence
*/
func (star *Star) GetSequence() string {
	return star.Sequence
}

/*
Returns the stars star classification
*/
func (star *Star) GetClass() string {
	return star.Class
}

/*
Returns the stars designation/name
*/
func (star *Star) GetName() string {
	return star.Name
}

/*
Returns the stars size in solar radii
*/
func (star *Star) GetSize() float64 {
	return star.SolarRadii
}

/*
Returns the stars mass in solar masses
*/
func (star *Star) GetMass() float64 {
	return star.SolarMasses
}

/*
Returns the stars mass in kg
*/
func (star *Star) GetMassKg() float64 {
	return star.GetMass() * consts.SOLAR_MASS
}

/*
Returns the stars volume in m^3
*/
func (star *Star) Volume() float64 {
	return (4.0 / 3.0) * math.Pi * math.Pow(star.SolarRadii*consts.SOLAR_RADII, 3)
}

/*
Returns the stars Density in g/cm^3
*/
func (star *Star) Density() float64 {
	return star.GetMassKg() / star.Volume()
}

/*
Returns the stars surface area in m^2
*/
func (star *Star) Area() float64 {
	return 4 * math.Pi * math.Pow(star.SolarRadii*consts.SOLAR_RADII, 2)
}

/*
Returns the stars luminosity in solar luminosities
*/
func (star *Star) Luminosity() float64 {
	return (4 * math.Pi * math.Pow(star.SolarRadii*consts.SOLAR_RADII, 2) * consts.STEFAN_BOLTZMANN_CONST * math.Pow(star.TemperatureK, 4)) / consts.SOLAR_LUMINOSITY
}

/*
Returns the stars Escape Velocity in m/s
*/
func (star *Star) EscapeVelocity() (m float64) {
	return (math.Sqrt((2 * consts.GRAVITATIONAL_CONST * star.GetMassKg()) / (star.SolarRadii * consts.SOLAR_RADII)))
}

/*
Returns the stars Escape Velocity in m/s
*/
func (star *Star) SolarWindVelocity() float64 {
	return math.Sqrt((2 * consts.BOLTZMANN_CONST * star.TemperatureK) / consts.PROTON_MASS)
}

/*
Returns the stars solar wind mass loss rate in kg/s
*/
func (star *Star) SolarWindMassLossRate() float64 {
	return consts.SOLAR_WIND_MASS_LOSS_RATE * star.Area() * star.SolarWindVelocity()
}

/*
Returns the stars thermal/termination shock distance in AU
*/
func (star *Star) TerminalShockDistance() float64 {
	// Formula for termination shock distance (r in meters)
	r := math.Sqrt((star.SolarWindMassLossRate() * star.SolarWindVelocity()) / (4 * math.Pi * consts.ISM_DENSITY * math.Pow(consts.ISM_VELOCITY, 2)))

	// Convert distance from meters to AU
	return r / consts.AU
}

/*
Returns the stars frost line, in AU from the star
*/
func (star *Star) FrostLine() float64 {
	return 2.7 * math.Sqrt(star.Luminosity())
}

/*
Adjustments for the heat of a sun, when calculating Habitable Zone
Not scientifically  accurate, but gives a better result IMO
*/
func (star *Star) calculateLuminosityFactor() (IRFactor, UVFactor float64) {
	IRFactor = 1.0
	UVFactor = 1.0

	low := map[string]bool{
		"White Dwarf": true,
		"Red Dwarf":   true,
		"Brown Dwarf": true,
	}

	high := map[string]bool{
		"Supergiant":   true,
		"Bright Giant": true,
	}

	if _, found := low[star.Sequence]; found {
		IRFactor = 1 - consts.INFRRED_FACTOR*(consts.SOL_TEMP_K-star.TemperatureK)/consts.SOL_TEMP_K
	}

	if _, found := high[star.Sequence]; found {
		UVFactor = 1 + consts.UV_FACTOR*(star.TemperatureK-consts.SOL_TEMP_K)/consts.SOL_TEMP_K
	}

	return IRFactor, UVFactor
}

/*
Returns habitable zone as start and end distance in AU from the star
*/
func (star *Star) HabitableZone() (start, end float64) {
	// high luminosity stars have high radiation output, pushing the start of the habitable zone further out.
	// low luminosity stars have most of their heat in infrared, narrowing the habitable zone entirely
	irFactor, uvFactor := star.calculateLuminosityFactor()

	dInner := irFactor * math.Sqrt(star.Luminosity()/consts.INNER_EDGE_STELLAR_FLUX) * uvFactor
	dOuter := irFactor * math.Sqrt(star.Luminosity()/consts.OUTER_EDGE_STELLAR_FLUX)
	return dInner, dOuter
}
