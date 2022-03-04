package userHandlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	util "whm-api/utils"
)

func (h *handler) List(ctx *gin.Context) {
	users, err := h.controller.List()

	if err == http.StatusText(http.StatusOK) {
		util.APIResponse(ctx, "Successfully fetched Users!", http.StatusOK, http.MethodGet, users)
	} else {
		util.APIResponse(ctx, err, http.StatusInternalServerError, http.MethodGet, nil)
	}
}
