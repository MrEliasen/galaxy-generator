package internal

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mreliasen/ihniwiad/pkg/galaxy"
	"github.com/mreliasen/ihniwiad/pkg/utils"
)

type Args struct {
	Seed int64 `form:"seed"`
}

func EndpointNeighbourhood(c *gin.Context) {
	gSeed, err := strconv.ParseInt(c.Param("seed"), 10, 64)
	if err != nil {
		gSeed = time.Now().UnixMicro()
	}

	var args Args

	if c.ShouldBind(&args) != nil {
		args.Seed = time.Now().UnixMicro()
	}

	if args.Seed == 0 {
		args.Seed = time.Now().UnixMicro()
	}

	log.Println(gSeed, args.Seed)

	rng := utils.NewSeededRNG(gSeed)
	g := galaxy.New(rng, gSeed)
	n := g.GenerateStellarNeighbourhood(args.Seed)

	c.JSON(200, n)
}
