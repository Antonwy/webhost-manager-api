package docker

import (
	"github.com/docker/docker/api/types"
)

func (config *Config) ListNetworks() ([]types.NetworkResource, error) {
	networks, err := config.Client.NetworkList(config.Context, types.NetworkListOptions{})
	if err != nil {
		return []types.NetworkResource{}, err
	}

	return networks, nil
}
