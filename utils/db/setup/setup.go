package dbSetup

import (
	"whm-api/utils/db"
)

func InitSchema() {
	db.DB.MustExec(initialSchema)
}
