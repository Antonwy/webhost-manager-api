package db_containers

import (
	"fmt"
	"whm-api/utils/db"
)

func (container *DBContainer) Remove() error {
	_, err := db.DB.NamedExec("delete from containers where id = %s;", container.ID)

	if err != nil {
		fmt.Printf("Failed deleting container: %s because: %s\n", container, err)
		return err
	}

	return nil
}
