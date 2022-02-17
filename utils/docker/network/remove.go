package network

import (
	"fmt"
	"whm-api/utils/docker"
)

func Remove(id string, config docker.Config) error {
	if err := config.Client.NetworkRemove(config.Context, id); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
