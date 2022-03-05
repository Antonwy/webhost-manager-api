package zoneHandlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	util "whm-api/utils"
)

func (h *handler) Get(ctx *gin.Context) {
	id := ctx.Param("id")

	zone, err := h.controller.Get(id)

	if err == http.StatusText(http.StatusOK) {
		util.APIResponse(ctx, "Successfully fetched zones!", http.StatusOK, http.MethodGet, zone)
	} else {
		util.APIResponse(ctx, err, http.StatusInternalServerError, http.MethodGet, nil)
	}
}
