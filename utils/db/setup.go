package db

import (
	"fmt"
	util "whm-api/utils"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const initialSchema string = `
	create table if not exists stacks
	(
			id   varchar not null
					constraint stacks_pk
							primary key,
			name varchar not null
	);

	alter table stacks
			owner to whm;

	create unique index if not exists stacks_id_uindex
			on stacks (id);

	create unique index if not exists stacks_name_uindex
			on stacks (name);

	create table if not exists containers
	(
			stack_id varchar not null
					constraint containers_stacks_id_fk
							references stacks
							on delete cascade,
			id       varchar not null
					constraint container_pk
							primary key
	);

	alter table containers
			owner to whm;

	create unique index if not exists containers_id_uindex
			on containers (id);

	create table if not exists ports
	(
			port         integer not null,
			container_id varchar
					constraint ports_containers_id_fk
							references containers
							on delete cascade
	);

	alter table ports
			owner to whm;

	create unique index if not exists ports_port_uindex
			on ports (port);
`

var DB *sqlx.DB

func Setup() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		util.GodotEnv("POSTGRES_DB"),
		5432,
		util.GodotEnv("POSTGRES_USER"),
		util.GodotEnv("POSTGRES_PASSWORD"),
		util.GodotEnv("POSTGRES_DB"),
	)

	DB = sqlx.MustConnect("postgres", psqlInfo)

	DB.MustExec(initialSchema)
}
