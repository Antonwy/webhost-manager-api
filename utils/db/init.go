package db

import (
	"fmt"
	util "whm-api/utils"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

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
}
