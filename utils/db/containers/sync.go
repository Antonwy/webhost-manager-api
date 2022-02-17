package db_containers

import (
	"github.com/docker/docker/api/types"
)

func Sync(dockerContainers []types.Container) {
	dbContainers, err := ListContainers()
	if err != nil {
		return
	}

	for _, dbContainer := range dbContainers {
		if !containsIDDocker(dockerContainers, dbContainer.ID) {
			dbContainer.Remove()
		}
	}
}

func containsIDDocker(c []types.Container, id string) bool {
	for _, v := range c {
		if v.ID == id {
			return true
		}
	}

	return false
}
