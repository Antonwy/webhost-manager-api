package network

import (
	"fmt"
	"whm-api/utils/docker"

	"github.com/docker/docker/api/types"
)

func Create(name string, config docker.Config) (string, error) {
	res, err := config.Client.NetworkCreate(config.Context, name, types.NetworkCreate{})

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return res.ID, nil
}
