package createZoneHandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	createZoneController "whm-api/controllers/zone-controllers/create"

	util "whm-api/utils"
)

func (h *handler) Create(ctx *gin.Context) {
	var input createZoneController.InputCreateZone

	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.APIResponse(ctx, "Invalid zone input data: "+err.Error(), http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	zone, err := h.controller.Create(input)
	if err == http.StatusText(http.StatusOK) {
		util.APIResponse(ctx, "Successfully created WordPress Site!", http.StatusOK, http.MethodPost, zone)
	} else {
		util.APIResponse(ctx, err, http.StatusInternalServerError, http.MethodPost, nil)
	}
}
