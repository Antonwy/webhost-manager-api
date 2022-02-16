package ports

import (
	"fmt"
	"whm-api/utils/db"
)

func (port *Port) Create() error {
	_, err := db.DB.NamedExec("INSERT INTO ports (port, container_id) VALUES (:port, :container_id)", port)

	if err != nil {
		fmt.Printf("Failed inserting ports: %s because: %s\n", port, err)
		return err
	}

	return nil
}
