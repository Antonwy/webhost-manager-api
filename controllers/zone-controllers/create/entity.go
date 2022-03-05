package createZoneController

type InputCreateZone struct {
	ID   string `db:"id" json:"id" binding:"required"`
	Name string `db:"name" json:"name" binding:"required"`
}
