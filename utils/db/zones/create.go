package zones

import (
	"log"
	"whm-api/utils/db"
)

const insertQuery string = "INSERT INTO zones (id, name) VALUES (:id, :name)"

func (zone Zone) Create() error {
	_, err := db.DB.NamedExec(insertQuery, zone)

	if err != nil {
		log.Printf("Failed creating zone: %s because: %s\n", zone, err)
		return err
	}

	return nil
}
