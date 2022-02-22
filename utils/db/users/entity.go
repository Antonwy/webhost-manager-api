package users

type User struct {
	ID           string `json:"id" db:"id"`
	Email        string `json:"email" db:"email"`
	PasswordHash string `json:"password" db:"password_hash"`
	Role         string `json:"role" db:"role"`
}

func (user User) ToJWT() JWTUserClaims {
	return JWTUserClaims{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}
}

type JWTUserClaims struct {
	ID    string `json:"id" db:"id"`
	Email string `json:"email" db:"email"`
	Role  string `json:"role" db:"role"`
}
