package docker

import (
	"fmt"

	"github.com/docker/docker/api/types"
)

type NetworkID string

func CreateNetwork(name string, config Config) (string, error) {
	res, err := config.Client.NetworkCreate(config.Context, name, types.NetworkCreate{})

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return res.ID, nil
}
