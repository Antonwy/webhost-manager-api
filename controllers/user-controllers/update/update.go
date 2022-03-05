package updateUserController

import (
	"log"
	"net/http"
	"whm-api/utils/db/users"
)

func (c controller) Update(input InputUpdateUser) (users.User, string) {

	user, err := users.GetFromID(input.ID)

	if err != nil {
		log.Println(err)
		return users.User{}, "Couldn't find user with ID: " + input.ID
	}

	if input.Name != "" {
		user.Name = input.Name
	}

	if input.Email != "" {
		user.Email = input.Email
	}

	if input.Role != "" {
		user.Role = input.Role
	}

	if err := user.Update(); err != nil {
		log.Println(err)
		return users.User{}, "Couldn't update user!"
	}

	return user, http.StatusText(http.StatusOK)
}
