package internal

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mreliasen/ihniwiad/pkg/galaxy"
	"github.com/mreliasen/ihniwiad/pkg/utils"
)

type GalaxyArgs struct {
	Seed int64 `form:"seed"`
}

func EndpointGalaxy(c *gin.Context) {
	var args GalaxyArgs

	if c.ShouldBind(&args) != nil {
		args.Seed = time.Now().UnixMicro()
	}

	gSeed, err := strconv.ParseInt(c.Param("seed"), 10, 64)
	if err == nil {
		args.Seed = gSeed
	}

	if args.Seed == 0 {
		args.Seed = time.Now().UnixMicro()
	}

	rng := utils.NewSeededRNG(args.Seed)
	g := galaxy.New(rng, args.Seed)
	g.GenerateStellarNeighbourhood(rng.Int63())

	c.JSON(200, g)
}
