package zoneHandlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	util "whm-api/utils"
)

func (h *handler) Remove(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.controller.Remove(id); err == http.StatusText(http.StatusOK) {
		util.APIResponse(ctx, "Successfully deleted zone!", http.StatusOK, http.MethodDelete, nil)
	} else {
		util.APIResponse(ctx, err, http.StatusInternalServerError, http.MethodDelete, nil)
	}
}
