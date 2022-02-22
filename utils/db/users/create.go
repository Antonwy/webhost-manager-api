package users

import (
	"fmt"
	"whm-api/utils/db"
)

const insertQuery string = "INSERT INTO users (id, email, password_hash, role) VALUES (:id, :email, :password_hash, :role)"

func (user *User) Create() error {
	_, err := db.DB.NamedExec(insertQuery, user)

	if err != nil {
		panic(err)
		fmt.Printf("Failed inserting user: %s because: %s\n", user, err)
		return err
	}

	return nil
}
