package dockerContainer

import (
	"fmt"

	"github.com/docker/docker/api/types"
)

func (container DockerContainer) Start() error {
	if err := container.Config.Client.ContainerStart(container.Config.Context, container.ID, types.ContainerStartOptions{}); err != nil {
		fmt.Println(err)
		return err
	}

	dbContainer := container.AsDBContainer()
	dbContainer.Create()

	return nil
}
