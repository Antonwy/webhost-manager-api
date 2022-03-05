package createZoneHandler

import (
	createZoneController "whm-api/controllers/zone-controllers/create"
)

type handler struct {
	controller createZoneController.Controller
}

func NewHandler(controller createZoneController.Controller) *handler {
	return &handler{
		controller: controller,
	}
}
