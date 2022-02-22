package dbSetup

import (
	"fmt"
	util "whm-api/utils"
	passwordService "whm-api/utils/auth/password"
	"whm-api/utils/db"
	"whm-api/utils/db/users"

	"github.com/docker/distribution/uuid"
)

func AddUserAndSchema() {
	db.DB.MustExec(initialSchema)

	insertAdmin()
}

func insertAdmin() {
	fmt.Println("Inserting Admin...")
	email := util.GodotEnv("USER_EMAIL")
	password := util.GodotEnv("USER_PASSWORD")

	fmt.Printf("Email: %s, Password: %s\n", email, password)

	var dbUser users.User
	if err := db.DB.Get(&dbUser, "select * from users where email = $1;", email); err == nil {
		fmt.Println("Admin already exists!")
		return
	}

	fmt.Println("No Admin")

	passwordHash, err := passwordService.HashPassword(password)

	fmt.Println("Password hash: " + passwordHash)

	if err != nil {
		panic("Couldn't generate a hash for the password!")
	}

	admin := users.User{
		ID:           uuid.Generate().String(),
		Email:        email,
		PasswordHash: passwordHash,
		Role:         "Admin",
	}

	if err := admin.Create(); err != nil {
		panic("Couldn't create admin user!")
	}
}
