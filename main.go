package main

import (
	"log"
	"net/http"
	"whm-api/middleware"
	route "whm-api/routes"
	util "whm-api/utils"
	"whm-api/utils/auth"
	"whm-api/utils/db"
	dbSetup "whm-api/utils/db/setup"

	"github.com/docker/docker/client"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/supertokens/supertokens-golang/supertokens"
)

// example: https://github.com/restuwahyu13/gin-rest-api
// another: https://dev.to/techschoolguru/implement-restful-http-api-in-go-using-gin-4ap1

func main() {

	db.Setup()
	dbSetup.InitSchema()
	auth.Setup()
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

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000", "http://127.0.0.1:3000", "https://antonwy.me"},
		AllowMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowHeaders: append([]string{"content-type"},
			supertokens.GetAllCORSHeaders()...),
		AllowCredentials: true,
	}))

	router.Use(gzip.Gzip(gzip.BestCompression))

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		panic(err)
	}

	router.Use(func(c *gin.Context) {
		supertokens.Middleware(http.HandlerFunc(
			func(rw http.ResponseWriter, r *http.Request) {
				c.Next()
			})).ServeHTTP(c.Writer, c.Request)

		c.Abort()
	})

	router.Use(middleware.VerifySession(nil))
	router.Use(middleware.ExtractSessionData)

	apiRouter := router.Group("/v1")

	route.InitDockerRoutes(apiRouter, cli)
	route.InitWordPressRoutes(apiRouter, cli)
	route.InitStackRoutes(apiRouter, cli)
	route.InitCloudFlareRoutes(apiRouter)
	route.InitUserRoutes(apiRouter)
	route.InitZoneRoutes(apiRouter)

	return router
}
