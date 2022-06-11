package createStackHandler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	createStackController "whm-api/controllers/stacks-controllers/create"
	util "whm-api/utils"
)

type handler struct {
	controller createStackController.Controller
}

func NewHandler(controller createStackController.Controller) *handler {
	return &handler{
		controller: controller,
	}
}

func (h *handler) CreateStackHandler(ctx *gin.Context) {
	var input createStackController.CreateStackInput

	bindErr := ctx.ShouldBindJSON(&input)
	if bindErr != nil {
		util.APIResponse(ctx, bindErr.Error(), http.StatusInternalServerError, http.MethodPost, nil)
		return
	}

	res, err := h.controller.CreateStack(input)

	if err == http.StatusText(http.StatusOK) {
		util.APIResponse(ctx, "Successfully created Stacks!", http.StatusOK, http.MethodPost, res.Response())
	} else {
		fmt.Println(err)
		util.APIResponse(ctx, err, http.StatusInternalServerError, http.MethodPost, nil)
	}
}
