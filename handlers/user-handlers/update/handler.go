package updateUserHandler

import (
	updateUserController "whm-api/controllers/user-controllers/update"
)

type handler struct {
	controller updateUserController.Controller
}

func NewHandler(controller updateUserController.Controller) *handler {
	return &handler{
		controller: controller,
	}
}
