package dockerContainer

import (
	"errors"
)

func (container *DockerContainer) QuickStart() error {
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
