package dockerContainer

import (
	"errors"
	"strconv"
	db_containers "whm-api/utils/db/containers"

	"github.com/docker/docker/api/types"
	"github.com/docker/go-connections/nat"
)

func (c *DockerContainer) Validate() error {
	dockerContainers, err := c.Config.ListContainers()
	if err != nil {
		return err
	}

	db_containers.Sync(dockerContainers)

	for _, container := range dockerContainers {
		if equalContainerNames(container.Names, c.Name) {
			return errors.New("There is already a container with the name " + c.Name)
		}

		if err := exposesSamePort(container.Ports, c.PortBindings); err != nil {
			return err
		}
	}

	if c.NetworkName != "" {
		dockerNetworks, err := c.Config.ListNetworks()

		if err != nil {
			return err
		}

		for _, network := range dockerNetworks {
			if network.Name == c.NetworkName {
				return errors.New("There is already a network with the name " + c.NetworkName)
			}
		}
	}

	return nil
}

func exposesSamePort(ports []types.Port, containerPorts map[nat.Port]nat.Port) error {
	for _, port := range ports {
		strToPort := strconv.Itoa(int(port.PublicPort))
		natToPort := nat.Port(strToPort)
		if _, ok := containerPorts[natToPort]; ok {
			return errors.New("There is already a container, that exposes port " + strToPort)
		}
	}

	return nil
}

func equalContainerNames(cNames []string, name string) bool {

	for _, cName := range cNames {
		cuttedName := cName[1:]

		if cuttedName == name {
			return true
		}
	}

	return false
}
