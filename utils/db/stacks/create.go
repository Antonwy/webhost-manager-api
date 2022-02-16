package stacks

import (
	"fmt"
	"whm-api/utils/db"

	"github.com/docker/distribution/uuid"
)

func (stack *Stack) Create() error {
	_, err := db.DB.NamedExec("INSERT INTO stacks (id, name) VALUES (:id, :name)", stack)

	if err != nil {
		fmt.Printf("Failed inserting stack: %s because: %s\n", stack, err)
		return err
	}

	return nil
}

func Create(name string) (Stack, error) {

	stack := Stack{
		Name: name,
		ID:   uuid.Generate().String(),
	}

	_, err := db.DB.NamedExec("INSERT INTO stacks (id, name) VALUES (:id, :name)", stack)

	if err != nil {
		fmt.Printf("Failed inserting stack: %s because: %s\n", stack, err)
		return Stack{}, err
	}

	return stack, nil
}
