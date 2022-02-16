package db_containers

import (
	"fmt"
	"whm-api/utils/db"
)

func (container *DBContainer) Create() error {
	_, err := db.DB.NamedExec("INSERT INTO containers (id, stack_id) VALUES (:id, :stack_id)", container)

	if err != nil {
		fmt.Printf("Failed inserting container: %s because: %s\n", container, err)
		return err
	}

	return nil
}
