package internal

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mreliasen/ihniwiad/pkg/utils"
)

var rng = utils.NewSeededRNG(time.Now().UnixMicro())

func EndpointDocs(c *gin.Context) {
	headers := []string{
		"Method", "Endpoints", "Response Type", "Args", "Notes",
	}
	data := [][]string{
		{"GET", "/api/galaxy", "JSON", "?seed=<number> (optional, galaxy seed)", "Generates a galaxy with a single stellar neighbourhood"},
		{"GET", "/api/galaxy/<galaxy seed>/neighbourhood", "JSON", "?seed=<number> (optional, neighbourhood seed)", "Generates another stellar neighbourhood for the specified galaxy"},
	}

	table := utils.ToTable(headers, data)

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"table": strings.Join(table, "\n"),
	})
}
