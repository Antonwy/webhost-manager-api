package db_containers

import (
	"database/sql"
)

type DBContainer struct {
	ID      string         `db:"id"`
	StackID sql.NullString `db:"stack_id"`
	Ports   []string
}
