package db_containers

import (
	"whm-api/utils/db"
)

func ListContainers() ([]DBContainer, error) {
	dbContainers := []DBContainer{}
	err := db.DB.Select(&dbContainers, "select * from containers;")

	if err != nil {
		return nil, err
	}

	return dbContainers, nil
}

func ListContainersFromStack(stackID string) ([]DBContainer, error) {
	dbContainers := []DBContainer{}
	err := db.DB.Select(&dbContainers, "select * from containers where stack_id = $1;", stackID)

	if err != nil {
		return nil, err
	}

	return dbContainers, nil
}
