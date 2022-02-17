package removeStackHandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	removeStack "whm-api/controllers/stacks-controllers/remove"
	util "whm-api/utils"
)

type handler struct {
	service removeStack.Service
}

func NewHandler(service removeStack.Service) *handler {
	return &handler{service: service}
}

func (h *handler) RemoveStackHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.service.RemoveStackService(id)

	if err == http.StatusText(http.StatusOK) {
		util.APIResponse(ctx, "Successfully removed Stack!", http.StatusOK, http.MethodDelete, nil)
	} else {
		util.APIResponse(ctx, err, http.StatusInternalServerError, http.MethodDelete, nil)
	}
}
