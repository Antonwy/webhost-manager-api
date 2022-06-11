package templateControllers

import "whm-api/utils/templates"

type Controller interface {
	ListTemplates() ([]templates.TemplateResponse, string)
}

type controller struct{}

func NewController() *controller {
	return &controller{}
}
