package route

import (
	"github.com/gin-gonic/gin"

	listContainers "whm-api/controllers/docker-controllers/container-controllers/list"
	listContainersHandler "whm-api/handlers/docker-handlers/container-handlers/list"

	createWordPress "whm-api/controllers/docker-controllers/wordpress-controllers/create"
	createWordPressHandler "whm-api/handlers/docker-handlers/wordpress-handlers/create"

	"github.com/docker/docker/client"
)

func InitDockerRoutes(router *gin.Engine) {

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	// DOCKER
	listContainersRepository := listContainers.NewRepositoryCreate(cli)
	listContainersService := listContainers.NewServiceCreate(listContainersRepository)
	listContainersHandler := listContainersHandler.NewHandlerListContainers(listContainersService)

	containersRoute := router.Group("/api/v1/docker") //.Use(middleware.Auth())
	containersRoute.GET("/containers", listContainersHandler.ListContainersHandler)

	// WORDPRESS
	wordpressRoute := containersRoute.Group("/wordpress")
	InitDockerWordPressRoutes(wordpressRoute, cli)
}

func InitDockerWordPressRoutes(router *gin.RouterGroup, cli *client.Client) {
	createWordPressRepository := createWordPress.NewRepositoryCreate(cli)
	createWordPressService := createWordPress.NewServiceCreate(createWordPressRepository)
	createWordPressHandler := createWordPressHandler.NewHandlerCreateWordPress(createWordPressService)

	router.POST("/", createWordPressHandler.CreateWordPressHandler)
}
