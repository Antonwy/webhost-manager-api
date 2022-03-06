package handlerCloudflare

import (
	"github.com/gin-gonic/gin"
	"net/http"
	util "whm-api/utils"
)

func (h *handler) DeleteRecordHandler(ctx *gin.Context) {
	zoneID := ctx.Param("zoneId")
	recordID := ctx.Param("recordId")

	status := h.controller.DeleteRecordController(zoneID, recordID)

	if status == http.StatusText(http.StatusOK) {
		util.APIResponse(ctx, "Successfully deleted record", http.StatusOK, http.MethodDelete, nil)
	} else {
		util.APIResponse(ctx, "Failed deleting record with error: "+status, http.StatusInternalServerError, http.MethodDelete, nil)
	}
}
