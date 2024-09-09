package internal

import (
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/mreliasen/ihniwiad/pkg/galaxy/public/interfaces"
)

func genScatter3dData(n interfaces.StellarNeighbourhoodInterface) []opts.Chart3DData {
	data := make([]opts.Chart3DData, 0)

	for _, s := range n.GetSystems() {
		star := s.GetStar()

		data = append(data, opts.Chart3DData{
			Name: star.GetName(),
			Value: []interface{}{
				star.GetCoordinate().GetX(),
				star.GetCoordinate().GetY(),
				star.GetCoordinate().GetZ(),
			},
			ItemStyle: &opts.ItemStyle{
				Color: star.GetColour(),
			},
		})
	}
	return data
}

func scatter3DBase(n interfaces.StellarNeighbourhoodInterface) *charts.Scatter3D {
	scatter3d := charts.NewScatter3D()
	scatter3d.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Neighbourhood Stellar Mass Objects (Stars)",
		}),
		charts.WithXAxis3DOpts(opts.XAxis3D{Name: "LY", Show: opts.Bool(true)}),
		charts.WithYAxis3DOpts(opts.YAxis3D{Name: "LY"}),
		charts.WithZAxis3DOpts(opts.ZAxis3D{Name: "LY"}),
	)

	scatter3d.AddSeries("Stars/Star Systems", genScatter3dData(n))
	return scatter3d
}

func PlotExample(n interfaces.StellarNeighbourhoodInterface) {
	page := components.NewPage()
	page.AddCharts(
		scatter3DBase(n),
	)

	f, err := os.Create("scatter3d.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
