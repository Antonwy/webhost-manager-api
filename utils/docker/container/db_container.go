package dockerContainer

import (
	"database/sql"
	db_containers "whm-api/utils/db/containers"
)

func (container *DockerContainer) AsDBContainer() db_containers.DBContainer {
	ports := []string{}

	for port := range container.PortBindings {
		ports = append(ports, string(port))
	}

	return db_containers.DBContainer{
		ID: container.ID,
		StackID: sql.NullString{
			String: container.StackID,
			Valid:  true,
		},
		Ports: ports,
	}
}
