package utils

import (
	"math"
	"math/rand"

	"github.com/mreliasen/ihniwiad/pkg/coordinate"
	"github.com/mreliasen/ihniwiad/pkg/coordinate/public/interfaces"
	"github.com/mreliasen/ihniwiad/pkg/galaxy/public/consts"
)

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func RandomCartesianCoord(rng *rand.Rand, minDist, maxDist float64) interfaces.CoordinateInterface {
	theta := rng.Float64() * 2 * math.Pi

	u := rng.Float64()
	r := math.Sqrt(u*(maxDist*maxDist-minDist*minDist) + minDist*minDist)

	// Convert polar coordinates to Cartesian coordinates
	x := r * math.Cos(theta)
	y := r * math.Sin(theta)

	return coordinate.New(x, y)
}

func RandomPointInCylindricalAnnulus(rng *rand.Rand, r1, r2, z1, z2 float64) interfaces.CoordinateInterface {
	// Step 1: Random radius with uniform area distribution
	u := rng.Float64()
	r := math.Sqrt(u*(r2*r2-r1*r1) + r1*r1)

	// Step 2: Random angle theta
	theta := rng.Float64() * 2 * math.Pi

	// Step 3: Random z coordinate
	v := rng.Float64()
	z := v*(z2-z1) + z1

	// Convert polar coordinates to Cartesian coordinates
	x := r * math.Cos(theta)
	y := r * math.Sin(theta)

	c := coordinate.New(x, y)
	c.SetZ(z)
	return c
}

func CartesianToPolar(x, y float64) (radii float64, theta float64) {
	radii = math.Sqrt(x*x + y*y)
	theta = math.Atan2(y, x)

	return radii, theta
}

/*
The Hill Sphere represents the region around a smaller body (such as a planet) where its gravitational influence
dominates over the gravitational influence of a larger body (such as a star).
This region defines the space where satellites can orbit the smaller body without being pulled away by the larger one.

	a:  Semi-Major axis of the smaller body (eg planet or moon), in AU
	m:  Mass of the smaller body (eg planet or moon), in kg
	mm: Mass of the major body (eg. planet or star) in kg

	Returns the limit in AU
*/
func HillSphere(a, m, mm float64) float64 {
	return ((a * math.Pow(m/(3*mm), 1.0/3.0)) / 1000) / consts.AU
}

/*
The Roche Limit is the minimum distance at which a celestial object, like a planet or a moon, can approach a larger
body, like a star or planet, without being torn apart by tidal forces.

	r:  Larger body radii in meters
	d:  Larger body density in kg/m^3
	dd: Smaller body density in kg/m^3

	Returns the limit in AU
*/
func RocheLimit(r, d, dd float64) float64 {
	return ((r * math.Pow(2*(d/dd), 1.0/3.0)) / 1000) / consts.AU
}

func CelciusToKelvin(c float64) float64 {
	return c + 273.15
}

func KelvinToCelcius(k float64) float64 {
	return k - 273.15
}
