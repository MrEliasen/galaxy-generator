package internal

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/mreliasen/ihniwiad/pkg/galaxy/public/interfaces"
	"github.com/mreliasen/ihniwiad/pkg/utils"
)

type Galaxy struct {
	Rng            *rand.Rand                                 `json:"-"`
	Type           string                                     `json:"type"`
	Thickness      float64                                    `json:"thickness"`
	Radius         float64                                    `json:"radii"`
	BulgeRadius    float64                                    `json:"bulge_radii"`
	Seed           int64                                      `json:"seed"`
	Neighbourhoods []interfaces.StellarNeighbourhoodInterface `json:"stellar_neighbourhoods"`
}

/*
Sets the seeded
*/
func (g *Galaxy) SetSeed(seed int64) interfaces.GalaxyInterface {
	g.Seed = seed
	return g
}

/*
Sets the seeded RNG
*/
func (g *Galaxy) SetRNG(rng *rand.Rand) interfaces.GalaxyInterface {
	g.Rng = rng
	return g
}

/*
Sets size of the galaxy bulge, by its radius.

The radius is in LY
*/
func (g *Galaxy) SetBulgeRadius(r float64) interfaces.GalaxyInterface {
	g.BulgeRadius = r
	return g
}

/*
Sets size of the galaxy, by its radius.

The radius is in LY
*/
func (g *Galaxy) SetRadius(r float64) interfaces.GalaxyInterface {
	g.Radius = r
	return g
}

/*
Sets disk Thickness
*/
func (g *Galaxy) SetThickness(t float64) interfaces.GalaxyInterface {
	g.Thickness = t
	return g
}

/*
Returns the "habitable zone" of a galaxy.

This is a contested subject, let it serve as no more than a guide.

Returns the habitable zone by its inner and outer distance to the galactic centre.
*/
func (g *Galaxy) HabitableZone() (inner float64, outer float64) {
	inner = utils.RoundFloat(g.Radius*0.18, 0) // ly
	outer = utils.RoundFloat(g.Radius*0.66, 0) // ly
	return inner, outer
}

/*
Generates the location of the stellar neighbourhood of a solar system within the "galactic habitable zone".
*/
func (g *Galaxy) GenerateStellarNeighbourhood(seed int64) interfaces.StellarNeighbourhoodInterface {
	hzInner, hzOuter := g.HabitableZone()

	nhRng := utils.NewSeededRNG(g.Seed ^ seed)

	dist := utils.RoundFloat(nhRng.Float64()*((hzOuter-hzInner)/100)+hzInner, 5)
	rad := utils.RoundFloat(float64(nhRng.Intn(30-15)+15), 0)                 // 15 LY min, up to 30 LY
	sDensity := utils.RoundFloat(nhRng.Float64()*((0.006-0.003)/10)+0.003, 5) // 0.003 min density, up to 0.006

	location := utils.RandomCartesianCoord(nhRng, dist, g.BulgeRadius+60, hzOuter)
	location.SetZ(nhRng.Float64()*(g.Thickness-rad) + rad)

	neighbourhood := &StellarNeighbourhood{
		Rng:     nhRng,
		Seed:    seed,
		Dist:    dist,
		Radius:  rad,
		Density: sDensity,
		Coords:  location,
	}

	neighbourhood.PopulateNeighbourhood()
	g.Neighbourhoods = append(g.Neighbourhoods, neighbourhood)

	return neighbourhood
}

func (g *Galaxy) Display() {
	g.dumpStarSystems()
	g.plotStarSystems()
}

func (g *Galaxy) plotStarSystems() {
	PlotExample(g.Neighbourhoods[0])
}

func (g *Galaxy) dumpStarSystems() {
	re := lipgloss.NewRenderer(os.Stdout)
	baseStyle := re.NewStyle().Padding(0, 1)
	headerStyle := baseStyle.Foreground(lipgloss.Color("252")).Bold(true)
	selectedStyle := baseStyle.Background(lipgloss.Color("#494949"))
	typeColors := map[string]lipgloss.Color{
		"O": lipgloss.Color("#9BB0FF"),
		"B": lipgloss.Color("#A9BFFF"),
		"A": lipgloss.Color("#CAD7FF"),
		"F": lipgloss.Color("#F8F7FF"),
		"G": lipgloss.Color("#FFF4EA"),
		"K": lipgloss.Color("#FFD2A1"),
		"M": lipgloss.Color("#FFCC6F"),
	}
	dimTypeColors := map[string]lipgloss.Color{
		"O": lipgloss.Color("#9BB0FF"),
		"B": lipgloss.Color("#A9BFFF"),
		"A": lipgloss.Color("#CAD7FF"),
		"F": lipgloss.Color("#F8F7FF"),
		"G": lipgloss.Color("#FFF4EA"),
		"K": lipgloss.Color("#FFD2A1"),
		"M": lipgloss.Color("#FFCC6F"),
	}

	headers := []string{"Designation", "Sequence", "Spectral Class", "Luminosity Class", "Size (R☉)", "Mass (M☉)", "Stellar Coordinates"}
	data := [][]string{}

	solarStyle := lipgloss.NewStyle()

	for _, n := range g.Neighbourhoods {
		for _, s := range n.GetSystems() {
			data = append(data, []string{
				s.GetStar().GetName(),
				fmt.Sprintf("%s %s", solarStyle.Foreground(lipgloss.Color(s.GetStar().GetColour())).Render("█████"), s.GetStar().GetSequence()),
				s.GetStar().GetClass(),
				s.GetStar().GetLuminosityClass(),
				fmt.Sprintf("%0.2f", s.GetStar().GetSize()),
				fmt.Sprintf("%0.2f", s.GetStar().GetMass()),
				fmt.Sprintf("%.4f, %.4f, %.4f", s.GetStar().GetCoordinate().GetX(), s.GetStar().GetCoordinate().GetY(), s.GetStar().GetCoordinate().GetZ()),
			})
		}
	}

	t := table.New().
		Headers(headers...).
		Border(lipgloss.NormalBorder()).
		BorderStyle(re.NewStyle().Foreground(lipgloss.Color("238"))).
		Width(140).
		Rows(data...).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == 0 {
				return headerStyle
			}

			f, e := strconv.ParseFloat(data[row-1][4], 64)
			if e == nil && f >= 1 {
				return selectedStyle
			}

			even := row%2 == 0

			switch col {
			case 2, 3: // Type 1 + 2
				c := typeColors
				if even {
					c = dimTypeColors
				}

				color := c[fmt.Sprint(data[row-1][col])]
				return baseStyle.Foreground(color)
			}

			if even {
				return baseStyle.Foreground(lipgloss.Color("245"))
			}

			return baseStyle.Foreground(lipgloss.Color("252"))
		})

	fmt.Println(t)
}
