package users

import (
	"log"
	"whm-api/utils/db"
)

const insertQuery string = "INSERT INTO users (id, name, email, role) VALUES (:id, :name, :email, :role)"

func (user User) Create() error {
	_, err := db.DB.NamedExec(insertQuery, user)

	if err != nil {
		log.Printf("Failed creating user: %s because: %s\n", user, err)
		return err
	}

	return nil
}
