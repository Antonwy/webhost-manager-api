package dockerContainer

import (
	"whm-api/utils/docker"

	"github.com/docker/go-connections/nat"
)

type DockerContainer struct {
	Config       docker.Config         `json:"config"`
	ID           string                `json:"id"`
	Name         string                `json:"name"`
	Image        string                `json:"image"`
	Env          []string              `json:"env"`
	Volumes      map[string]struct{}   `json:"volumes"`
	NetworkID    string                `json:"network_id"`
	NetworkName  string                `json:"network_name"`
	PortBindings map[nat.Port]nat.Port `json:"port_bindings"`
	StackID      string                `json:"stack_id"`
}
