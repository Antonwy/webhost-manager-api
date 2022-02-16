package handlerWordPress

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	createWordPress "whm-api/controllers/docker-controllers/wordpress-controllers/create"
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
	ctx.ShouldBindJSON(&input)

	validate := validator.New()

	validateErr := validate.Struct(input)

	if validateErr != nil {
		util.APIResponse(ctx, validateErr.Error(), http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	res, err := h.service.CreateWordPressService(&input)

	if err == http.StatusText(http.StatusOK) {
		util.APIResponse(ctx, "Successfully created wordpress page", http.StatusOK, http.MethodGet, res)
	} else {
		util.APIResponse(ctx, "Failed creating wordpress with error: "+err, http.StatusInternalServerError, http.MethodGet, nil)
	}
}
