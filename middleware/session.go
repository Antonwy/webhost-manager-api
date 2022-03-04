package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"net/http"
	util "whm-api/utils"
	"whm-api/utils/constants"
)

func ExtractSessionData(c *gin.Context) {
	sessionContainer, err := session.GetSession(c.Request, c.Writer, nil)
	if err != nil {
		util.APIResponse(c, "Couldn't extract session from request!", http.StatusForbidden, c.Request.Method, nil)
		c.Abort()
		return
	}

	c.Set(constants.SessionUserIdKey, sessionContainer.GetUserID())
	c.Next()
}
