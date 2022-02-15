package handlerContainers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	listContainers "whm-api/controllers/docker-controllers/container-controllers/list"
	util "whm-api/utils"
)

type handler struct {
	service listContainers.Service
}

func NewHandlerListContainers(service listContainers.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ListContainersHandler(ctx *gin.Context) {
	containers, errListContainers := h.service.ListContainersService()

	if errListContainers != http.StatusText(http.StatusOK) {
		util.APIResponse(ctx, "Successfully fetched docker containers", http.StatusOK, http.MethodGet, containers)
	} else {
		util.APIResponse(ctx, "Failed fetching docker containers with error: "+errListContainers, http.StatusInternalServerError, http.MethodGet, containers)
	}
}
