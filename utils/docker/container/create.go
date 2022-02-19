package dockerContainer

import (
	"fmt"
	"whm-api/utils/docker"
	"whm-api/utils/docker/network"

	dcontainer "github.com/docker/docker/api/types/container"
	dockerNetwork "github.com/docker/docker/api/types/network"
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
		container.Config.Client.NetworkConnect(container.Config.Context, net, res.ID, &dockerNetwork.EndpointSettings{})
	}

	if container.ConnectToProxyNetwork {
		fmt.Println("Connecting with " + docker.ProxyNetworkName)
		id, err := network.IdFromName(docker.ProxyNetworkName, container.Config)
		if err != nil {
			fmt.Println(err)
			return err
		}

		if err := container.Config.Client.NetworkConnect(container.Config.Context, id, res.ID, &dockerNetwork.EndpointSettings{}); err != nil {
			fmt.Println(err)
			return err
		}
	}

	if err != nil {
		fmt.Println(err)
		return err
	}

	container.ID = res.ID

	return nil
}
