package listStackHandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	listStacks "whm-api/controllers/stacks-controllers/list"
	util "whm-api/utils"
)

type handler struct {
	service listStacks.Service
}

func NewHandler(service listStacks.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ListStacksHandler(ctx *gin.Context) {
	res, err := h.service.ListStacksService()

	if err == http.StatusText(http.StatusOK) {
		util.APIResponse(ctx, "Successfully fetched Stacks!", http.StatusOK, http.MethodGet, res)
	} else {
		util.APIResponse(ctx, err, http.StatusInternalServerError, http.MethodGet, nil)
	}
}
