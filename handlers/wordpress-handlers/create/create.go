package createWordPressHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	createWordPress "whm-api/controllers/wordpress-controllers/create"
	util "whm-api/utils"
)

type handler struct {
	service createWordPress.Service
}

func NewHandlerCreateWordPress(service createWordPress.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateWordPressHandler(ctx *gin.Context) {
	var input createWordPress.InputCreateWordPress
	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.APIResponse(ctx, err.Error(), http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	res, err := h.service.CreateWordPressService(&input)

	if err == http.StatusText(http.StatusOK) {
		util.APIResponse(ctx, "Successfully created WordPress Site!", http.StatusOK, http.MethodGet, res)
	} else {
		util.APIResponse(ctx, err, http.StatusInternalServerError, http.MethodGet, nil)
	}
}
