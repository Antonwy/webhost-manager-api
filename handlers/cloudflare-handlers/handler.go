package handlerCloudflare

import cloudflareControllers "whm-api/controllers/cloudflare-controllers"

type handler struct {
	controller cloudflareControllers.Controller
}

func NewHandler(controller cloudflareControllers.Controller) *handler {
	return &handler{
		controller: controller,
	}
}
