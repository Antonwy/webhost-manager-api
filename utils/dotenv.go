package util

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GodotEnv(key string) string {
	env := make(chan string, 1)

	if os.Getenv("GO_ENV") == "production" {
		godotenv.Load(".env")
		env <- os.Getenv(key)
	} else {
		godotenv.Load(".env.local")
		env <- os.Getenv(key)
	}

	return <-env
}

func GodotEnvBool(key string, defaultVal bool) bool {
	b, err := strconv.ParseBool(GodotEnv(key))

	if err != nil {
		return defaultVal
	} else {
		return b
	}
}
