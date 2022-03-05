package users

import (
	"log"
	"whm-api/utils/db"
)

const updateQuery string = "update users set name = :name, role = :role where id = :id"

func (user User) Update() error {
	_, err := db.DB.NamedExec(updateQuery, user)

	if err != nil {
		log.Printf("Failed updating user: %s because: %s\n", user, err)
		return err
	}

	return nil
}
