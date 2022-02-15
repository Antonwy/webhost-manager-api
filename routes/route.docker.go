package route

import (
	"github.com/gin-gonic/gin"

	listContainers "whm-api/controllers/docker-controllers/container-controllers/list"
	listContainersHandler "whm-api/handlers/docker-handlers/container-handlers/list"

	"github.com/docker/docker/client"
)

func InitDockerRoutes(router *gin.Engine) {

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	listContainersRepository := listContainers.NewRepositoryCreate(cli)
	listContainersService := listContainers.NewServiceCreate(listContainersRepository)
	listContainersHandler := listContainersHandler.NewHandlerListContainers(listContainersService)

	groupRoute := router.Group("/api/v1/docker") //.Use(middleware.Auth())
	groupRoute.GET("/containers", listContainersHandler.ListContainersHandler)
}
