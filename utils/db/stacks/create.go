package stacks

import (
	"log"
	"os"
	"whm-api/utils/db"
	"whm-api/utils/docker"

	"github.com/docker/distribution/uuid"
)

const insertQuery string = "INSERT INTO stacks (id, name, type, network_name, url) VALUES (:id, :name, :type, :network_name, :url)"

func (stack *Stack) Create() error {
	_, err := db.DB.NamedExec(insertQuery, stack)

	if err != nil {
		log.Printf("Failed inserting stack: %s because: %s\n", stack, err)
		return err
	}

	dirName := stack.DirectoryName()

	if err := os.MkdirAll("/data/stacks/"+dirName, os.ModePerm); err != nil {
		log.Printf("Failed creating directory %s because: %s\n", dirName, err)
		return err
	}

	return nil
}

func Create(config docker.Config, name string) (Stack, error) {

	stack := Stack{
		Name:   name,
		ID:     GenerateStackID(),
		Config: config,
	}

	stack.Create()

	return stack, nil
}

func GenerateStackID() string {
	return uuid.Generate().String()
}
