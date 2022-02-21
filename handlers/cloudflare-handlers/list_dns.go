package handlerCloudflare

import (
	"net/http"

	"github.com/gin-gonic/gin"

	util "whm-api/utils"
)

func (h *handler) ListDNSHandler(ctx *gin.Context) {
	zones, status := h.controller.ListDNSController(ctx.Param("id"))

	if status == http.StatusText(http.StatusOK) {
		util.APIResponse(ctx, "Successfully fetched zone dns", http.StatusOK, http.MethodGet, zones)
	} else {
		util.APIResponse(ctx, "Failed listing zone dns with error: "+status, http.StatusInternalServerError, http.MethodGet, nil)
	}
}
