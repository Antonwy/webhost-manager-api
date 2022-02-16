package dockerContainer

import (
	"whm-api/utils/docker"

	"github.com/docker/go-connections/nat"
)

type DockerContainer struct {
	Config       docker.Config
	ID           string
	Name         string
	Image        string
	Env          []string
	Volumes      map[string]struct{}
	NetworkID    string
	NetworkName  string
	PortBindings map[nat.Port]nat.Port
	StackID      string
}
