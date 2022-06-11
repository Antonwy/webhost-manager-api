package templateHandlers

import (
	applicationControllers "whm-api/controllers/template-controllers"
)

type handler struct {
	controller applicationControllers.Controller
}

func NewHandler(controller applicationControllers.Controller) *handler {
	return &handler{
		controller: controller,
	}
}
