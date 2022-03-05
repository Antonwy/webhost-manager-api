package updateUserController

type InputUpdateUser struct {
	ID    string `db:"id" json:"id"`
	Name  string `db:"name" json:"name" binding:"omitempty,min=4,max=15"`
	Email string `db:"email" json:"email" binding:"omitempty,email"`
	Role  string `db:"role" json:"role"`
}
