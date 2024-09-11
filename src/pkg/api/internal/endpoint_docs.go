package internal

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mreliasen/ihniwiad/pkg/utils"
)

var rng = utils.NewSeededRNG(time.Now().UnixMicro())

func EndpointDocs(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}
