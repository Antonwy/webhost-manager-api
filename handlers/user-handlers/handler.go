package userHandlers

import (
	userControllers "whm-api/controllers/user-controllers"
)

type handler struct {
	controller userControllers.Controller
}

func NewHandler(controller userControllers.Controller) *handler {
	return &handler{
		controller: controller,
	}
}
