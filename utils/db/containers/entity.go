package db_containers

type DBContainer struct {
	ID      string `db:"id"`
	StackID string `db:"stack_id"`
}
