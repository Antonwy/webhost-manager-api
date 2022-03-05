package zoneHandlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	util "whm-api/utils"
)

func (h *handler) List(ctx *gin.Context) {
	zones, err := h.controller.List()

	if err == http.StatusText(http.StatusOK) {
		util.APIResponse(ctx, "Successfully fetched zones!", http.StatusOK, http.MethodGet, zones)
	} else {
		util.APIResponse(ctx, err, http.StatusInternalServerError, http.MethodGet, nil)
	}
}
