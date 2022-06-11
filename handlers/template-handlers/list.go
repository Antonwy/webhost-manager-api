package templateHandlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	util "whm-api/utils"
)

func (h *handler) List(ctx *gin.Context) {
	templates, status := h.controller.ListTemplates()

	if status == http.StatusText(http.StatusOK) {
		util.APIResponse(ctx, "Successfully fetched templates", http.StatusOK, http.MethodGet, templates)
	} else {
		util.APIResponse(ctx, "Failed listing templates with error: "+status, http.StatusInternalServerError, http.MethodGet, nil)
	}
}
