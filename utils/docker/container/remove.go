package dockerContainer

import (
	"fmt"
	db_containers "whm-api/utils/db/containers"

	"github.com/docker/docker/api/types"
)

func (container *DockerContainer) Remove() {
	cli := container.Config.Client
	ctx := container.Config.Context

	if container.NetworkID != "" {
		err := cli.NetworkRemove(ctx, container.NetworkID)

		if err != nil {
			fmt.Println(err)
		}
	}

	if err := cli.ContainerRemove(ctx, container.ID, types.ContainerRemoveOptions{
		RemoveVolumes: true,
	}); err == nil {
		dbContainer := db_containers.DBContainer{
			ID:      container.ID,
			StackID: container.StackID,
		}

		dbContainer.Remove()
	}
}
