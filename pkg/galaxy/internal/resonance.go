package internal

import (
	"math"
	"math/rand"

	"github.com/mreliasen/ihniwiad/pkg/galaxy/public/interfaces"
	"github.com/mreliasen/ihniwiad/pkg/logger"
)

type Resonance struct {
	inner, outer int
}

var innerResonances = []Resonance{
	{6, 5},
	{5, 4},
	{4, 3},
	{3, 2},
	{2, 1},
}

var midResonances = []Resonance{
	{4, 3},
	{7, 5},
	{5, 3},
	{3, 1},
}

var outerResonances = []Resonance{
	{8, 5},
	{7, 4},
	{9, 4},
	{7, 3},
	{5, 2},
}

// Titius-Bode distribuion
func getTitiusBodeDistribution(rng *rand.Rand, star interfaces.StarInterface) []float64 {
	starAdjustment := star.FrostLine() / 2.7
	semiMajorAxis := 0.4 * starAdjustment
	minDistance := semiMajorAxis / 4

	solarScaleFactor := star.SolarWindVelocity() / 1_000_000
	initialSemiMajorAxis := rng.Float64()*(semiMajorAxis-minDistance) + minDistance
	growthFactor := 2.0 + solarScaleFactor // Factor by which the distance increases

	distances := []float64{
		initialSemiMajorAxis,
	}

	for i := 1; i < 30; i++ {
		d := initialSemiMajorAxis * math.Pow(growthFactor, float64(i))
		distances = append(distances, d)
	}

	return distances
}

// Random function to select a resonance from the list
func getRandomResonance(rng *rand.Rand, resonances []Resonance) Resonance {
	return resonances[rng.Intn(len(resonances))]
}

// Function to calculate the next planet's semi-major axis based on resonance
func getResonanceFactor(resonance Resonance) float64 {
	resonanceFactor := math.Pow(float64(resonance.inner)/float64(resonance.outer), 2.0/3.0)
	return resonanceFactor
}

// Function to generate planets with resonances
func generatePlanetDistances(rng *rand.Rand, system interfaces.StarSystemInterface) (inner []float64, outer []float64) {
	logger := logger.Get()

	// Store the semi-major axes of the planets
	tbDist := getTitiusBodeDistribution(rng, system.GetStar())
	planetDistances := make([]float64, len(tbDist))

	// cap rocky planets based on how many "inner" there are
	frostLine := system.GetStar().FrostLine()
	hzInner, hzOuter := system.GetStar().HabitableZone()

	for i := 0; i < len(tbDist)-1; i++ {
		if i == 0 {
			planetDistances[0] = tbDist[0]
			inner = append(inner, tbDist[0])
			continue
		}

		// Determine the zone based on the previous planet's distance
		var resonance Resonance
		if planetDistances[i-1] < hzInner {
			// Inner zone
			resonance = getRandomResonance(rng, innerResonances)
		} else if planetDistances[i-1] >= hzInner && planetDistances[i-1] <= hzOuter {
			// Middle zone
			resonance = getRandomResonance(rng, midResonances)
		} else if planetDistances[i-1] > hzOuter && planetDistances[i-1] < frostLine {
			// Middle zone
			resonance = getRandomResonance(rng, midResonances)
		} else {
			// Outer zone
			resonance = getRandomResonance(rng, outerResonances)
		}

		dist := tbDist[i-1] + getResonanceFactor(resonance)

		// limit it to 15k AU, just to limit it somewhere
		if dist > 15_000 {
			break
		}

		if dist < frostLine {
			inner = append(inner, dist)
		} else {
			outer = append(outer, dist)
		}

		// Calculate the next planet's semi-major axis
		planetDistances[i] = dist
	}

	// Output the planet data
	for i, distance := range planetDistances {
		if distance <= 0.0 {
			break
		}

		logger.Debugf("Planet %d: Semi-major axis = %.2f AU\n", i+1, distance)
	}

	return inner, outer
}
