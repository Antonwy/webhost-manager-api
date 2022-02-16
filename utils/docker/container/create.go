package dockerContainer

import (
	"fmt"

	dcontainer "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
)

func (container *DockerContainer) Create() error {

	portBinding := nat.PortMap{}

	if container.PortBindings != nil {
		for k, v := range container.PortBindings {
			portBinding[v] = []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: string(k),
				},
			}
		}
	}

	res, err := container.Config.Client.ContainerCreate(container.Config.Context, &dcontainer.Config{
		Image:   container.Image,
		Env:     container.Env,
		Volumes: container.Volumes,
	}, &dcontainer.HostConfig{
		PortBindings: portBinding,
	}, nil, nil, container.Name)

	if net := container.NetworkID; net != "" {
		container.Config.Client.NetworkConnect(container.Config.Context, net, res.ID, &network.EndpointSettings{})
	}

	if err != nil {
		fmt.Println(err)
		return err
	}

	container.ID = res.ID

	return nil
}
