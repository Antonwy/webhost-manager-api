package userControllers

import (
	"fmt"
	"net/http"
	"whm-api/utils/db/users"
)

func (c controller) List() ([]users.User, string) {
	users, err := users.List()

	if err != nil {
		fmt.Println(err)
		return nil, "Couldn't get users"
	}

	return users, http.StatusText(http.StatusOK)
}
