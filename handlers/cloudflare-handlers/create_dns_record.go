package handlerCloudflare

import (
	"net/http"

	"github.com/gin-gonic/gin"

	util "whm-api/utils"
)

func (h *handler) CreateDNSRecordHandler(ctx *gin.Context) {
	zones, status := h.controller.CreateDNSRecordController(ctx.Param("id"), ctx.Request.Body)

	if status == http.StatusText(http.StatusOK) {
		util.APIResponse(ctx, "Successfully created DNS record!", http.StatusOK, http.MethodGet, zones)
	} else {
		util.APIResponse(ctx, "Failed creating DNS record with error: "+status, http.StatusInternalServerError, http.MethodGet, nil)
	}
}
