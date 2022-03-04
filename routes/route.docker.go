package route

import (
	"github.com/gin-gonic/gin"
	userControllers "whm-api/controllers/user-controllers"
	userHandlers "whm-api/handlers/user-handlers"

	listContainers "whm-api/controllers/docker-controllers/container-controllers/list"
	listStacks "whm-api/controllers/stacks-controllers/list"
	listStacksHandler "whm-api/handlers/stacks-handlers/list"

	removeStack "whm-api/controllers/stacks-controllers/remove"
	removeStackHandler "whm-api/handlers/stacks-handlers/remove"

	listContainersHandler "whm-api/handlers/docker-handlers/container-handlers/list"

	createWordPress "whm-api/controllers/wordpress-controllers/create"
	createWordPressHandler "whm-api/handlers/wordpress-handlers/create"

	cloudflareControllers "whm-api/controllers/cloudflare-controllers"
	handlerCloudflare "whm-api/handlers/cloudflare-handlers"

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

	router.POST("/wordpress", createWordPressHandler.CreateWordPressHandler)
}

func InitStackRoutes(router *gin.RouterGroup, cli *client.Client) {
	listStacksRepository := listStacks.NewRepository(cli)
	listStacksService := listStacks.NewService(listStacksRepository)
	listStacksHandler := listStacksHandler.NewHandler(listStacksService)

	removeStackRepository := removeStack.NewRepository(cli)
	removeStackService := removeStack.NewService(removeStackRepository)
	removeStackHandler := removeStackHandler.NewHandler(removeStackService)

	router.GET("/stacks", listStacksHandler.ListStacksHandler)
	router.DELETE("/stacks/:id", removeStackHandler.RemoveStackHandler)
}

func InitCloudFlareRoutes(router *gin.RouterGroup) {
	cloudflareController := cloudflareControllers.NewController()
	cloudflareHandler := handlerCloudflare.NewHandler(cloudflareController)

	router.GET("/zones", cloudflareHandler.ListZonesHandler)
	router.GET("/zones/:id/records", cloudflareHandler.ListDNSHandler)
	router.POST("/zones/:id/records", cloudflareHandler.CreateDNSRecordHandler)
}

func InitUserRoutes(router *gin.RouterGroup) {
	userController := userControllers.NewController()
	userHandler := userHandlers.NewHandler(userController)

	router.GET("/users/:id", userHandler.Get)
	router.GET("/users", userHandler.List)
}
