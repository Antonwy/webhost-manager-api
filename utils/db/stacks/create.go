package stacks

import (
	"fmt"
	"whm-api/utils/db"
	"whm-api/utils/docker"

	"github.com/docker/distribution/uuid"
)

const insertQuery string = "INSERT INTO stacks (id, name, type, network_name, url) VALUES (:id, :name, :type, :network_name, :url)"

func (stack *Stack) Create() error {
	_, err := db.DB.NamedExec(insertQuery, stack)

	if err != nil {
		fmt.Printf("Failed inserting stack: %s because: %s\n", stack, err)
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

	_, err := db.DB.NamedExec(insertQuery, stack)

	if err != nil {
		fmt.Printf("Failed inserting stack: %s because: %s\n", stack, err)
		return Stack{}, err
	}

	return stack, nil
}

func GenerateStackID() string {
	return uuid.Generate().String()
}

func CreateWithNetwork(config docker.Config, name string, networkName string) (Stack, error) {
	stack, err := Create(config, name)

	if err != nil {
		return Stack{}, err
	}

	stack.NetworkName = networkName

	return stack, nil
}
