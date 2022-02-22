package authController

import (
	inputLogin "whm-api/utils/auth/input-login"
)

type Controller interface {
	LoginController(input inputLogin.InputLogin) (string, string)
}

type controller struct{}

func NewController() *controller {
	return &controller{}
}
