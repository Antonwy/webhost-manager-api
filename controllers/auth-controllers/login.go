package authController

import (
	"net/http"
	inputLogin "whm-api/utils/auth/input-login"
	jwtService "whm-api/utils/auth/jwt"
	passwordService "whm-api/utils/auth/password"
	"whm-api/utils/db/users"
)

func (c controller) LoginController(input inputLogin.InputLogin) (string, string) {
	user, err := users.GetFromEmail(input.Email)

	if err != nil {
		return "", "Couldn't find a user with email: " + input.Email
	}

	if !passwordService.CheckPasswordHash(input.Password, user.PasswordHash) {
		return "", "Password not correct!"
	}

	service := jwtService.JWTAuthService()

	return service.GenerateToken(user), http.StatusText(http.StatusOK)
}
