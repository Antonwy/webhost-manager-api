package route

import (
	"github.com/gin-gonic/gin"
	userControllers "whm-api/controllers/user-controllers"
	updateUserController "whm-api/controllers/user-controllers/update"
	zoneControllers "whm-api/controllers/zone-controllers"
	createZoneController "whm-api/controllers/zone-controllers/create"
	userHandlers "whm-api/handlers/user-handlers"
	updateUserHandler "whm-api/handlers/user-handlers/update"
	zoneHandlers "whm-api/handlers/zone-handlers"
	createZoneHandler "whm-api/handlers/zone-handlers/create"

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

	cloudflareApi := router.Group("/cloudflare")
	cloudflareApi.GET("/zones", cloudflareHandler.ListZonesHandler)
	cloudflareApi.GET("/zones/:id/records", cloudflareHandler.ListDNSHandler)
	cloudflareApi.POST("/zones/:id/records", cloudflareHandler.CreateDNSRecordHandler)
}

func InitUserRoutes(router *gin.RouterGroup) {
	userController := userControllers.NewController()
	userHandler := userHandlers.NewHandler(userController)
	updateUserC := updateUserController.NewController()
	updateUserH := updateUserHandler.NewHandler(updateUserC)

	router.GET("/users/me", userHandler.GetMe)
	router.GET("/users/:id", userHandler.Get)
	router.PUT("/users/:id", updateUserH.Update)
	router.GET("/users", userHandler.List)
}

func InitZoneRoutes(router *gin.RouterGroup) {
	zoneController := zoneControllers.NewController()
	zoneHandler := zoneHandlers.NewHandler(zoneController)
	createZoneC := createZoneController.NewController()
	createZoneH := createZoneHandler.NewHandler(createZoneC)

	router.GET("/zones", zoneHandler.List)
	router.GET("/zones/:id", zoneHandler.Get)
	router.DELETE("/zones/:id", zoneHandler.Remove)
	router.POST("/zones", createZoneH.Create)
	router.POST("/zones/sync", zoneHandler.Sync)
}
