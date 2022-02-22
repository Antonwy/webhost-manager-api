package handlerAuth

import (
	authController "whm-api/controllers/auth-controllers"
)

type handler struct {
	controller authController.Controller
}

func NewHandler(controller authController.Controller) *handler {
	return &handler{
		controller: controller,
	}
}
