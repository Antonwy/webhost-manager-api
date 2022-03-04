package userControllers

import "whm-api/utils/db/users"

type Controller interface {
	Get(id string) (users.User, string)
	List() ([]users.User, string)
}

type controller struct{}

func NewController() *controller {
	return &controller{}
}
