package users

import "whm-api/utils/db"

func List() ([]User, error) {
	var users []User

	if err := db.DB.Select(&users, "select * from users"); err != nil {
		return nil, err
	}

	return users, nil
}
