package handlerCloudflare

import (
	"net/http"

	"github.com/gin-gonic/gin"

	util "whm-api/utils"
)

func (h *handler) CreateDNSRecordHandler(ctx *gin.Context) {
	zone, status := h.controller.CreateDNSRecordController(ctx.Param("id"), ctx.Request.Body)

	if status == http.StatusText(http.StatusOK) {
		util.APIResponse(ctx, "Successfully created DNS record!", http.StatusOK, http.MethodPost, zone)
	} else {
		util.APIResponse(ctx, status, http.StatusInternalServerError, http.MethodPost, nil)
	}
}
