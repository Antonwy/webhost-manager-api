package userHandlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	util "whm-api/utils"
	"whm-api/utils/constants"
)

func (h *handler) GetMe(ctx *gin.Context) {
	id, ok := ctx.Get(constants.SessionUserIdKey)

	if !ok {
		util.APIResponse(ctx, "Session userid couldn't be found in context!", http.StatusInternalServerError, http.MethodGet, nil)
	}

	user, err := h.controller.Get(id.(string))

	if err == http.StatusText(http.StatusOK) {
		util.APIResponse(ctx, "Successfully fetched User!", http.StatusOK, http.MethodGet, user)
	} else {
		util.APIResponse(ctx, err, http.StatusInternalServerError, http.MethodGet, nil)
	}
}
