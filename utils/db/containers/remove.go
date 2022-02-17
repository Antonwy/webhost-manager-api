package db_containers

import (
	"fmt"
	"whm-api/utils/db"
	"whm-api/utils/docker"

	"github.com/docker/docker/api/types"
)

func (container *DBContainer) Remove() error {
	_, err := db.DB.NamedExec("delete from containers where id = :id;", container)

	if err != nil {
		fmt.Printf("Failed deleting container: %s because: %s\n", container, err)
		return err
	}

	return nil
}

func (container *DBContainer) RemoveWithContainer(config docker.Config) error {

	if err := config.Client.ContainerRemove(config.Context, container.ID, types.ContainerRemoveOptions{
		RemoveVolumes: true,
		Force:         true,
	}); err != nil {
		fmt.Printf("Error removing container %s because: $s\n", container, err)
		return err
	}

	_, err := db.DB.NamedExec("delete from containers where id = :id;", container)

	if err != nil {
		fmt.Printf("Failed deleting container: %s because: %s\n", container, err)
		return err
	}

	return nil
}
