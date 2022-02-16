package createWordPress

type InputCreateWordPress struct {
	Name       string `json:"name" validate:"required"`
	WPUsername string `json:"wp_username" validate:"required"`
	WPPassword string `json:"wp_password" validate:"required"`
	DBUsername string `json:"db_username" validate:"required"`
	DBPassword string `json:"db_password" validate:"required"`
}
