package dockerContainer

import (
	"errors"
	"whm-api/utils/docker"
)

func (container *DockerContainer) QuickStart() error {
	if networkName := container.NetworkName; networkName != "" {
		id, err := docker.CreateNetwork(networkName, container.Config)

		if err == nil {
			container.NetworkID = id
		}
	}

	if err := container.PullImage(); err != nil {
		return errors.New("Error pulling " + container.Image)
	}

	if err := container.Create(); err != nil {
		return errors.New("Error creating " + container.Image)
	}

	if err := container.Start(); err != nil {
		return errors.New("Error starting " + container.Image + " container")
	}

	return nil
}
