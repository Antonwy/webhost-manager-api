package users

import "whm-api/utils/db"

func GetFromID(id string) (User, error) {
	user := User{}
	if err := db.DB.Get(&user, "select * from users where id = $1", id); err != nil {
		return User{}, err
	}

	return user, nil
}

func GetFromEmail(email string) (User, error) {
	user := User{}
	if err := db.DB.Get(&user, "select * from users where email = $1", email); err != nil {
		return User{}, err
	}

	return user, nil
}
