package docker

import "github.com/docker/docker/api/types"

func (config *Config) ListContainers() ([]types.Container, error) {
	containers, err := config.Client.ContainerList(config.Context, types.ContainerListOptions{})
	if err != nil {
		return []types.Container{}, err
	}

	return containers, nil
}
