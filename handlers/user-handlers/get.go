package userHandlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	util "whm-api/utils"
)

func (h *handler) Get(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := h.controller.Get(id)

	if err == http.StatusText(http.StatusOK) {
		util.APIResponse(ctx, "Successfully fetched User!", http.StatusOK, http.MethodGet, user)
	} else {
		util.APIResponse(ctx, err, http.StatusInternalServerError, http.MethodGet, nil)
	}
}
