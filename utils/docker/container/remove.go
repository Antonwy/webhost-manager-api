package dockerContainer

import (
	"fmt"

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
		Force:         true,
	}); err == nil {
		dbContainer := container.AsDBContainer()

		dbContainer.Remove()
	}
}
