package ports

type Port struct {
	ContainerID string `db:"container_id"`
	Port        string `db:"port"`
}
