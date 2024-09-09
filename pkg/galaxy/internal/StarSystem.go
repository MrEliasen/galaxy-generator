package internal

import (
	"math/rand"

	"github.com/mreliasen/ihniwiad/pkg/galaxy/internal/planet"
	"github.com/mreliasen/ihniwiad/pkg/galaxy/public/consts"
	"github.com/mreliasen/ihniwiad/pkg/galaxy/public/interfaces"
	"github.com/mreliasen/ihniwiad/pkg/logger"
	"github.com/mreliasen/ihniwiad/shared/utils"
)

type StarSystem struct {
	rng          *rand.Rand
	star         interfaces.StarInterface
	planets      []interfaces.PlanetInterface
	gasGiants    int
	iceGiants    int
	rockyPlanets int
}

func (s *StarSystem) GetPlanetTypeCount() (rocky, iceGiants, gasGiants, total int) {
	return s.rockyPlanets, s.iceGiants, s.gasGiants, s.rockyPlanets + s.iceGiants + s.gasGiants
}

func (s *StarSystem) SetIceGiants(n int) {
	s.iceGiants = n
}

func (s *StarSystem) SetGasGiants(n int) {
	s.gasGiants = n
}

func (s *StarSystem) SetRockyPlants(n int) {
	s.rockyPlanets = n
}

func (s *StarSystem) SetStar(st interfaces.StarInterface) interfaces.StarSystemInterface {
	s.star = st
	return s
}

func (s *StarSystem) SetRNG(r *rand.Rand) {
	s.rng = r
}

func (s *StarSystem) GetStar() interfaces.StarInterface {
	return s.star
}

func (s *StarSystem) dumpDetails() {
	logger := logger.Get()

	i, o := s.GetStar().HabitableZone()
	logger.Debugf("\nSystem: %s\n", s.star.GetName())
	logger.Debugf("Star Size (R☉): %.2f\n", s.GetStar().GetSize())
	logger.Debugf("Star Temp(K): %.2f\n", s.GetStar().GetTemperatureK())
	logger.Debugf("Star Luminosity(L☉): %.2f\n", s.GetStar().Luminosity())
	logger.Debugf("Frost Line(AU): %.2f\n", s.GetStar().FrostLine())
	logger.Debugf("Habitable Zone(AU): %.2f - %.2f\n", i, o)
}

func (s *StarSystem) PopulatePlanets() {
	logger := logger.Get()
	s.dumpDetails()

	numRockyPlanets, numIceGiants, numGasGiants, total := s.GetPlanetTypeCount()

	if total == 0 {
		return
	}

	innerAxis, outerAxis := generatePlanetDistances(s.rng, s)

	// inside snowline
	innerPlanets := 0
	outerPlanets := 0
	for _, d := range innerAxis {
		if d >= s.GetStar().FrostLine() {
			break
		}

		if innerPlanets >= numRockyPlanets {
			break
		}

		innerPlanets++
	}

	for range outerAxis {
		if outerPlanets >= numGasGiants+numIceGiants {
			break
		}

		outerPlanets++
	}

	planetsList := []interfaces.PlanetInterface{}

	starRadiiM := s.GetStar().GetSize() * consts.SOLAR_RADII
	starDensityKg := s.GetStar().Density() * 1000

	p := s.rng.Float64() + 1.0                                                       // rand between 1.0-2.0
	leftoverSolarMass := (s.GetStar().GetMass() / (100.0 - p)) * p                   // find the remaining % mass left for planets
	leftoverEarthMass := (leftoverSolarMass * consts.SOLAR_MASS) / consts.EARTH_MASS // as above, in earth masses
	hzInner, hzOuter := s.GetStar().HabitableZone()
	logger.Debugf("Leftover Mass Total: %0.4f\n", leftoverEarthMass)

	// generate inner planets, if they are not within the roche limit of the star, we skip it
	for d := range innerPlanets {
		if len(innerAxis) <= d {
			break
		}

		p := planet.GenerateRockyPlanet(s.rng)
		rl := utils.RocheLimit(starRadiiM, starDensityKg, p.GetDensity()*1000)

		if rl > innerAxis[d] {
			logger.Debugf("%0.4f > %0.4f\n", rl, innerAxis[d])
			continue
		}

		leftoverEarthMass -= p.GetMass() / consts.EARTH_MASS
		p.SetOrbitDistance(innerAxis[d])

		if p.GetOrbitDistance() > hzInner && p.GetOrbitDistance() < hzOuter {
			p.SetHabitable(true)
		}

		planetsList = append(planetsList, p)
	}

	logger.Debugf("Rocky Planets: %d\n", len(planetsList))

	// generate the outer planets, if they are within the roche limit of the star, we skip it
	generatedGasGiants := 0
	generatedIceGiants := 0
	for d := range outerPlanets {
		if len(outerAxis) <= d {
			break
		}

		var p interfaces.PlanetInterface

		if generatedGasGiants < numGasGiants {
			p = planet.GenerateGasGiant(s.rng)
			generatedGasGiants++
		} else if generatedIceGiants < numIceGiants {
			p = planet.GenerateIceGiant(s.rng)
			generatedIceGiants++
		} else {
			break
		}

		if p == nil {
			break
		}

		rl := utils.RocheLimit(starRadiiM, starDensityKg, p.GetDensity()*1000)

		if rl > outerAxis[d] {
			continue
		}

		leftoverEarthMass -= p.GetMass() / consts.EARTH_MASS
		p.SetOrbitDistance(outerAxis[d])
		planetsList = append(planetsList, p)
	}

	logger.Debugf("Gas/Ice Planets: %d/%d\n", generatedGasGiants, generatedIceGiants)

	fl := false
	hzI := false
	hzO := false
	for _, e := range planetsList {
		if !hzI && e.GetOrbitDistance() > hzInner {
			logger.Debugf("AU: %0.2f -> -------- (Habitable Zone Start)\n", hzInner)
			hzI = true
		}
		if !hzO && e.GetOrbitDistance() > hzOuter {
			logger.Debugf("AU: %0.2f -> -------- (Habitable Zone End)\n", hzOuter)
			hzO = true
		}
		if !fl && e.GetOrbitDistance() > s.GetStar().FrostLine() {
			logger.Debugf("AU: %0.2f -> -------- (frost_line)\n", s.GetStar().FrostLine())
			fl = true
		}
		logger.Debugf("AU: %0.2f -> %0.2f ER, %0.2f EM (%0.4f g/cm^3) (%s)\n", e.GetOrbitDistance(), e.GetSize(), e.GetMass()/consts.EARTH_MASS, e.GetDensity(), e.GetType())
	}

	logger.Debugf("Leftover Mass: %0.4f\n", leftoverEarthMass)

	s.planets = planetsList
}
