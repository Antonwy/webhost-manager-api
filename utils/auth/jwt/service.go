package jwtService

import (
	"fmt"
	"os"
	"time"
	"whm-api/utils/db/users"

	"github.com/golang-jwt/jwt"
)

//jwt service
type JWTService interface {
	GenerateToken(user users.User) string
	ValidateToken(token string) (*jwt.Token, error)
}

type authCustomClaims struct {
	users.JWTUserClaims
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    "whm-api",
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) GenerateToken(user users.User) string {
	claims := &authCustomClaims{
		user.ToJWT(),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token: %s", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}
