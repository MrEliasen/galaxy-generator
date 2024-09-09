package interfaces

import "math/rand"

type GalaxyInterface interface {
	SetRNG(rng *rand.Rand) GalaxyInterface
	/*
		Sets size of the galaxy, by its radius.

		The radius is in LY
	*/
	SetRadius(radii float64) GalaxyInterface
	/*
		Sets disk Thickness outside of the bulge

		Thickness in LY
	*/
	SetThickness(ly float64) GalaxyInterface

	/*
		Returns the "habitable zone" of a galaxy.

		This is a contested subject, let it serve as no more than a guide.

		Returns the habitable zone by its inner and outer distance to the galactic centre.
	*/
	HabitableZone() (inner float64, outer float64)

	/*
		Generates the location of the stellar neighbourhood of a solar system within the "galactic habitable zone".
	*/
	GenerateStellarNeighbourhood() StellarNeighbourhoodInterface

	/*
		Dumps information about the galaxy, stellar neighbourhoods and systems to console
	*/
	Display()
}
