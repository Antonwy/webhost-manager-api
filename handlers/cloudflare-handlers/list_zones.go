package handlerCloudflare

import (
	"net/http"

	"github.com/gin-gonic/gin"

	util "whm-api/utils"
)

func (h *handler) ListZonesHandler(ctx *gin.Context) {
	zones, status := h.controller.ListZonesController()

	if status == http.StatusText(http.StatusOK) {
		util.APIResponse(ctx, "Successfully fetched cloudflare zones", http.StatusOK, http.MethodGet, zones)
	} else {
		util.APIResponse(ctx, "Failed listing cloudflare zones with error: "+status, http.StatusInternalServerError, http.MethodGet, nil)
	}
}
