package zoneHandlers

import (
	zoneControllers "whm-api/controllers/zone-controllers"
)

type handler struct {
	controller zoneControllers.Controller
}

func NewHandler(controller zoneControllers.Controller) *handler {
	return &handler{
		controller: controller,
	}
}
