package zones

import (
	"log"
	"whm-api/utils/db"
)

const updateQuery string = "update zones set name = :name, synced_with_cloudflare = :synced_with_cloudflare where id = :id"

func (zone Zone) Update() error {
	_, err := db.DB.NamedExec(updateQuery, zone)

	if err != nil {
		log.Printf("Failed updating zone: %s because: %s\n", zone, err)
		return err
	}

	return nil
}
