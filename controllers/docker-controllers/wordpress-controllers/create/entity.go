package createWordPress

type InputCreateWordPress struct {
	Name       string `json:"name" validate:"required"`
	Port       string `json:"port" validate:"required"`
	DBUsername string `json:"db_username" validate:"required"`
	DBPassword string `json:"db_password" validate:"required"`
}
