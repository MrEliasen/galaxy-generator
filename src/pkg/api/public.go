package api

import (
	"embed"
	"html/template"

	g "github.com/CAFxX/httpcompression/contrib/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	"github.com/mreliasen/ihniwiad/pkg/api/internal"
)

//go:embed internal/templates/*
var templateFS embed.FS

func Run() {
	compressor, _ := g.DefaultAdapter()

	r := gin.Default()
	r.Use(compressor)

	tmpl := template.Must(template.New("").ParseFS(templateFS, "internal/templates/*"))
	r.SetHTMLTemplate(tmpl)

	r.GET("/api/galaxy", internal.EndpointGalaxy)
	r.GET("/api/galaxy/:seed/neighbourhood", internal.EndpointNeighbourhood)
	r.GET("/", internal.EndpointDocs)

	r.Run("0.0.0.0:8081")
}
