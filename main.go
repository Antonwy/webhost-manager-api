package main

import (
	"log"
	route "whm-api/routes"
	util "whm-api/utils"
	"whm-api/utils/db"

	"github.com/docker/docker/client"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// example: https://github.com/restuwahyu13/gin-rest-api
// another: https://dev.to/techschoolguru/implement-restful-http-api-in-go-using-gin-4ap1

func main() {

	db.Setup()
	router := SetupRouter()

	log.Fatal(router.Run(":" + util.GodotEnv("GO_PORT")))

}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	if util.GodotEnv("GO_ENV") != "production" && util.GodotEnv("GO_ENV") != "test" {
		gin.SetMode(gin.DebugMode)
	} else if util.GodotEnv("GO_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router.Use(cors.Default())

	router.Use(gzip.Gzip(gzip.BestCompression))

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		panic(err)
	}

	apiRouter := router.Group("/v1")
	route.InitDockerRoutes(apiRouter, cli)
	route.InitWordPressRoutes(apiRouter, cli)
	route.InitStackRoutes(apiRouter, cli)
	route.InitCloudFlareRoutes(apiRouter)

	return router
}
