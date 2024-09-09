package interfaces

import "math/rand"

type StarSystemInterface interface {
	/*
		Returns the target number of rocky planets, ice giants, gas giants and the toal planet count.

		This is used when populating the star system, and will be updated once population if done, in case it couldn't fit
		all the planets for example.
	*/
	GetPlanetTypeCount() (rocky, iceGiants, gasGiants, total int)
	/*
		This is used when populating the star system, and will be updated once population if done, in case it couldn't fit
		all the planets for example.
	*/
	SetIceGiants(n int)
	/*
		This is used when populating the star system, and will be updated once population if done, in case it couldn't fit
		all the planets for example.
	*/
	SetGasGiants(n int)
	/*
		This is used when populating the star system, and will be updated once population if done, in case it couldn't fit
		all the planets for example.
	*/
	SetRockyPlants(n int)

	SetRNG(r *rand.Rand)
	/*
		Add a star to the star system, currently only supports single-star star systems.
	*/
	SetStar(StarInterface) StarSystemInterface
	GetStar() StarInterface
	/*
		Populates the star system with plants, based on the number of rocky, ice and gas giants set.
	*/
	PopulatePlanets()
}
