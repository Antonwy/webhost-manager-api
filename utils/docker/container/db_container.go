package dockerContainer

import db_containers "whm-api/utils/db/containers"

func (container *DockerContainer) AsDBContainer(stackID string) db_containers.DBContainer {
	return db_containers.DBContainer{
		ID:      container.ID,
		StackID: stackID,
	}
}
