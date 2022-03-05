package updateUserController

import "whm-api/utils/db/users"

type Controller interface {
	Update(InputUpdateUser) (users.User, string)
}

type controller struct{}

func NewController() *controller {
	return &controller{}
}
