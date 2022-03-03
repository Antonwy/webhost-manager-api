package userControllers

import (
	"fmt"
	"net/http"
	"whm-api/utils/db/users"
)

func (c controller) Get(id string) (users.User, string) {
	user, err := users.GetFromID(id)

	if err != nil {
		fmt.Println(err)
		return users.User{}, "Couldn't get user with id: " + id
	}

	return user, http.StatusText(http.StatusOK)
}
