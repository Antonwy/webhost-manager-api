package route

import (
	"github.com/gin-gonic/gin"

	listContainers "whm-api/controllers/docker-controllers/container-controllers/list"
	listStacks "whm-api/controllers/stacks-controllers/list"
	listStacksHandler "whm-api/handlers/stacks-handlers/list"

	removeStack "whm-api/controllers/stacks-controllers/remove"
	removeStackHandler "whm-api/handlers/stacks-handlers/remove"

	listContainersHandler "whm-api/handlers/docker-handlers/container-handlers/list"

	createWordPress "whm-api/controllers/wordpress-controllers/create"
	createWordPressHandler "whm-api/handlers/wordpress-handlers/create"

	"github.com/docker/docker/client"
)

func InitDockerRoutes(router *gin.RouterGroup, cli *client.Client) {
	// DOCKER
	listContainersRepository := listContainers.NewRepositoryCreate(cli)
	listContainersService := listContainers.NewServiceCreate(listContainersRepository)
	listContainersHandler := listContainersHandler.NewHandlerListContainers(listContainersService)

	containersRoute := router.Group("/docker") //.Use(middleware.Auth())
	containersRoute.GET("/containers", listContainersHandler.ListContainersHandler)
}

func InitWordPressRoutes(router *gin.RouterGroup, cli *client.Client) {
	createWordPressRepository := createWordPress.NewRepositoryCreate(cli)
	createWordPressService := createWordPress.NewServiceCreate(createWordPressRepository)
	createWordPressHandler := createWordPressHandler.NewHandlerCreateWordPress(createWordPressService)

	wordpressRoute := router.Group("/wordpress") //.Use(middleware.Auth())
	wordpressRoute.POST("/", createWordPressHandler.CreateWordPressHandler)
}

func InitStackRoutes(router *gin.RouterGroup, cli *client.Client) {
	listStacksRepository := listStacks.NewRepository(cli)
	listStacksService := listStacks.NewService(listStacksRepository)
	listStacksHandler := listStacksHandler.NewHandler(listStacksService)

	removeStackRepository := removeStack.NewRepository(cli)
	removeStackService := removeStack.NewService(removeStackRepository)
	removeStackHandler := removeStackHandler.NewHandler(removeStackService)

	stacksRoute := router.Group("/stacks") //.Use(middleware.Auth())
	stacksRoute.GET("/", listStacksHandler.ListStacksHandler)
	stacksRoute.DELETE("/:id", removeStackHandler.RemoveStackHandler)
}
