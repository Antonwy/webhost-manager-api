package userControllers

import "whm-api/utils/db/users"

type Controller interface {
	Get(id string) (users.User, string)
}

type controller struct{}

func NewController() *controller {
	return &controller{}
}
