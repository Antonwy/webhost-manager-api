package createWordPress

type InputCreateWordPress struct {
	Name       string `json:"name" binding:"required"`
	Url        string `json:"url" binding:"required"`
	DBUsername string `json:"db_username" binding:"required"`
	DBPassword string `json:"db_password" binding:"required"`
}
