package updateUserHandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	updateUserController "whm-api/controllers/user-controllers/update"
	util "whm-api/utils"
)

func (h *handler) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var input updateUserController.InputUpdateUser

	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.APIResponse(ctx, "Invalid user data: "+err.Error(), http.StatusBadRequest, http.MethodPut, nil)
		return
	}

	input.ID = id

	user, err := h.controller.Update(input)

	if err == http.StatusText(http.StatusOK) {
		util.APIResponse(ctx, "Successfully updated User!", http.StatusOK, http.MethodPut, user)
	} else {
		util.APIResponse(ctx, err, http.StatusInternalServerError, http.MethodPut, nil)
	}
}
