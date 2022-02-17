package network

import (
	"errors"
	"whm-api/utils/docker"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
)

func IdFromName(name string, config docker.Config) (string, error) {
	filter := filters.NewArgs()
	filter.Add("name", name)

	networks, err := config.Client.NetworkList(config.Context, types.NetworkListOptions{Filters: filter})

	if err != nil {
		return "", err
	}

	if len(networks) == 0 {
		return "", errors.New("Couldn't find any networks with name: " + name)
	}
	return networks[0].ID, nil
}
