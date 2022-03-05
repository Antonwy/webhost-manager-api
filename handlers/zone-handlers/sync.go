package zoneHandlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	util "whm-api/utils"
)

func (h *handler) Sync(ctx *gin.Context) {
	zones, err := h.controller.Sync()

	if err == http.StatusText(http.StatusOK) {
		util.APIResponse(ctx, "Successfully synced zones!", http.StatusOK, http.MethodGet, zones)
	} else {
		util.APIResponse(ctx, err, http.StatusInternalServerError, http.MethodGet, nil)
	}
}
