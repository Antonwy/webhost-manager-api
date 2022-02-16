package stacks

import (
	dockerContainer "whm-api/utils/docker/container"
)

type Stack struct {
	ID         string `db:"id"`
	Name       string `db:"name"`
	Containers []dockerContainer.DockerContainer
}
