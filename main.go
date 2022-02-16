package main

import (
	"log"
	route "whm-api/routes"
	util "whm-api/utils"
	"whm-api/utils/db"

	"github.com/gin-gonic/gin"
)

// example: https://github.com/restuwahyu13/gin-rest-api

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

	route.InitDockerRoutes(router)

	return router
}
