package dockerContainer

import (
	"fmt"
	db_containers "whm-api/utils/db/containers"
	"whm-api/utils/db/ports"

	"github.com/docker/docker/api/types"
)

func (container DockerContainer) Start() error {
	if err := container.Config.Client.ContainerStart(container.Config.Context, container.ID, types.ContainerStartOptions{}); err != nil {
		fmt.Println(err)
		return err
	}

	dbContainer := db_containers.DBContainer{
		ID:      container.ID,
		StackID: container.StackID,
	}

	dbContainer.Create()

	for k, _ := range container.PortBindings {
		port := ports.Port{
			ContainerID: container.ID,
			Port:        string(k),
		}

		port.Create()
	}

	return nil
}
