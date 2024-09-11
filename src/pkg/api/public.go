package api

import (
	"embed"
	"html/template"
	"net/http"

	g "github.com/CAFxX/httpcompression/contrib/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	"github.com/mreliasen/ihniwiad/pkg/api/internal"
)

//go:embed internal/templates/*
var templateFS embed.FS

//go:embed internal/assets/images/favicon.ico
var favicon []byte

func Run() {
	compressor, _ := g.DefaultAdapter()

	r := gin.Default()
	r.Use(compressor)

	tmpl := template.Must(template.New("").ParseFS(templateFS, "internal/templates/*"))
	r.SetHTMLTemplate(tmpl)

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.Data(http.StatusOK, "image/x-icon", favicon)
	})

	r.GET("/api/galaxy", internal.EndpointGalaxy)
	r.GET("/api/galaxy/:seed/neighbourhood", internal.EndpointNeighbourhood)
	r.GET("/", internal.EndpointDocs)

	r.Run("0.0.0.0:8081")
}
