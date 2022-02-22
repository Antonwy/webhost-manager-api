package handlerAuth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	util "whm-api/utils"
	inputLogin "whm-api/utils/auth/input-login"
)

func (h *handler) LoginHandler(ctx *gin.Context) {

	var input inputLogin.InputLogin
	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.APIResponse(ctx, err.Error(), http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	user, err := h.controller.LoginController(input)

	if err == http.StatusText(http.StatusOK) {
		util.APIResponse(ctx, "Successfully logged in!", http.StatusOK, http.MethodPost, user)
	} else {
		util.APIResponse(ctx, err, http.StatusInternalServerError, http.MethodPost, nil)
	}

}
