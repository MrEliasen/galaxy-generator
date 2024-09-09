package star

import "errors"

func calculateSpectralClass(star *Star) (*SpectralConfig, error) {
	solarLumen := star.Luminosity()

	for _, v := range SpectralClassifications[star.Class] {
		if solarLumen < v.LuminosityMin || solarLumen > v.LuminosityMax {
			continue
		}

		if v.TempKelvinMin != 0.0 && v.TempKelvinMax != 0.0 {
			if star.TemperatureK < v.TempKelvinMin || star.TemperatureK > v.TempKelvinMax {
				continue
			}
		}

		return &v, nil
	}

	return nil, errors.New("Unable to match class, skip")
}

type SpectralConfig struct {
	Name          string
	Type          string
	Color         string
	LuminosityMax float64
	LuminosityMin float64
	TempKelvinMax float64
	TempKelvinMin float64
}

var SpectralClassifications = map[string][]SpectralConfig{
	"O": {
		{
			Name:          "Ia/Ib",
			Type:          "Supergiant",
			Color:         "#97BCFE",
			LuminosityMax: 999999,
			LuminosityMin: 7500,
		},
		{
			Name:          "II",
			Type:          "Bright Giant",
			Color:         "#97BCFE",
			LuminosityMax: 7499,
			LuminosityMin: 3000,
		},
		{
			Name:          "III",
			Type:          "Giant",
			Color:         "#97BCFE",
			LuminosityMax: 2999,
			LuminosityMin: 0.1,
		},
		{
			Name:          "VII",
			Type:          "White Dwarf",
			Color:         "#97BCFE",
			LuminosityMax: 0.09,
			LuminosityMin: 0.0,
		},
	},
	"B": {
		{
			Name:          "Ia/Ib",
			Type:          "Supergiant",
			Color:         "#AFCBFE",
			LuminosityMax: 999999,
			LuminosityMin: 7500,
		},
		{
			Name:          "II",
			Type:          "Bright Giant",
			Color:         "#AFCBFE",
			LuminosityMax: 7500,
			LuminosityMin: 3000,
		},
		{
			Name:          "III",
			Type:          "Giant",
			Color:         "#AFCBFE",
			LuminosityMax: 2999,
			LuminosityMin: 0.01,
			TempKelvinMax: 12999,
			TempKelvinMin: 10000,
		},
		{
			Name:          "IV",
			Type:          "Subgiant",
			Color:         "#AFCBFE",
			LuminosityMax: 2999,
			LuminosityMin: 0.1,
			TempKelvinMax: 16999,
			TempKelvinMin: 13000,
		},
		{
			Name:          "V",
			Type:          "Main Sequence",
			Color:         "#AFCBFE",
			LuminosityMax: 2999,
			LuminosityMin: 0.1,
			TempKelvinMax: 30000,
			TempKelvinMin: 17000,
		},
		{
			Name:          "VII",
			Type:          "White Dwarf",
			Color:         "#AFCBFE",
			LuminosityMax: 0.09,
			LuminosityMin: 0,
		},
	},
	"A": {
		{
			Name:          "Ia/Ib",
			Type:          "Supergiant",
			Color:         "#D7E5FD",
			LuminosityMax: 999999,
			LuminosityMin: 5000,
		},
		{
			Name:          "II",
			Type:          "Bright Giant",
			Color:         "#D7E5FD",
			LuminosityMax: 4999,
			LuminosityMin: 750,
		},
		{
			Name:          "III",
			Type:          "Giant",
			Color:         "#D7E5FD",
			LuminosityMax: 749,
			LuminosityMin: 200,
		},
		{
			Name:          "IV",
			Type:          "Subgiant",
			Color:         "#D7E5FD",
			LuminosityMax: 199,
			LuminosityMin: 50,
		},
		{
			Name:          "V",
			Type:          "Main Sequence",
			Color:         "#D7E5FD",
			LuminosityMax: 49,
			LuminosityMin: 1,
		},
		{
			Name:          "VII",
			Type:          "White Dwarf",
			Color:         "#D7E5FD",
			LuminosityMax: 0.9,
			LuminosityMin: 0,
		},
	},
	"F": {
		{
			Name:          "Ia/Ib",
			Type:          "Supergiant",
			Color:         "#FBFAFE",
			LuminosityMax: 999999,
			LuminosityMin: 2000,
		},
		{
			Name:          "II",
			Type:          "Bright Giant",
			Color:         "#FBFAFE",
			LuminosityMax: 1999,
			LuminosityMin: 100,
		},
		{
			Name:          "III",
			Type:          "Giant",
			Color:         "#FBFAFE",
			LuminosityMax: 99,
			LuminosityMin: 40,
		},
		{
			Name:          "IV",
			Type:          "Subgiant",
			Color:         "#FBFAFE",
			LuminosityMax: 39,
			LuminosityMin: 8,
		},
		{
			Name:          "V",
			Type:          "Main Sequence",
			Color:         "#FBFAFE",
			LuminosityMax: 7,
			LuminosityMin: 0.1,
		},
		{
			Name:          "VII",
			Type:          "White Dwarf",
			Color:         "#FBFAFE",
			LuminosityMax: 0.09,
			LuminosityMin: 0,
		},
	},
	"G": {
		{
			Name:          "Ia/Ib",
			Type:          "Supergiant",
			Color:         "#FEF6EB",
			LuminosityMax: 999999,
			LuminosityMin: 2500,
		},
		{
			Name:          "II",
			Type:          "Bright Giant",
			Color:         "#FEF6EB",
			LuminosityMax: 2499,
			LuminosityMin: 100,
		},
		{
			Name:          "III",
			Type:          "Giant",
			Color:         "#FEF6EB",
			LuminosityMax: 99,
			LuminosityMin: 15,
		},
		{
			Name:          "IV",
			Type:          "Subgiant",
			Color:         "#FEF6EB",
			LuminosityMax: 14,
			LuminosityMin: 3,
		},
		{
			Name:          "V",
			Type:          "Main Sequence",
			Color:         "#FEF6EB",
			LuminosityMax: 2,
			LuminosityMin: 0.05,
		},
		{
			Name:          "VII",
			Type:          "White Dwarf",
			Color:         "#FEF6EB",
			LuminosityMax: 0.04,
			LuminosityMin: 0,
		},
	},
	"K": {
		{
			Name:          "Ia/Ib",
			Type:          "Supergiant",
			Color:         "#FEE2BC",
			LuminosityMax: 999999,
			LuminosityMin: 2500,
		},
		{
			Name:          "II",
			Type:          "Bright Giant",
			Color:         "#FEE2BC",
			LuminosityMax: 2499,
			LuminosityMin: 200,
		},
		{
			Name:          "III",
			Type:          "Giant",
			Color:         "#FEE2BC",
			LuminosityMax: 199,
			LuminosityMin: 15,
		},
		{
			Name:          "IV",
			Type:          "Subgiant",
			Color:         "#FEE2BC",
			LuminosityMax: 14,
			LuminosityMin: 1,
		},
		{
			Name:          "V",
			Type:          "Main Sequence",
			Color:         "#FEE2BC",
			LuminosityMax: 0.9,
			LuminosityMin: 0.001,
		},
		{
			Name:          "VII",
			Type:          "White Dwarf",
			Color:         "#FEE2BC",
			LuminosityMax: 0.0009,
			LuminosityMin: 0,
		},
	},
	"M": {
		{
			Name:          "Ia/Ib",
			Type:          "Supergiant",
			Color:         "#FFB766",
			LuminosityMax: 999999,
			LuminosityMin: 2500,
		},
		{
			Name:          "II",
			Type:          "Bright Giant",
			Color:         "#FFB766",
			LuminosityMax: 2499,
			LuminosityMin: 800,
		},
		{
			Name:          "III",
			Type:          "Giant",
			Color:         "#FFB766",
			LuminosityMax: 799,
			LuminosityMin: 10,
		},
		{
			Name:          "V",
			Type:          "Red Dwarf",
			Color:         "#F77200",
			LuminosityMax: 9.9,
			LuminosityMin: 0.001,
		},
		{
			Name:          "VI",
			Type:          "Brown Dwarf",
			LuminosityMax: 0.0009,
			Color:         "#BB0603",
			LuminosityMin: 0,
		},
	},
}
