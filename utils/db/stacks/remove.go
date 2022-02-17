package stacks

import (
	"fmt"
	"whm-api/utils/db"
)

// func Remove(id string) error {
// 	_, err := db.DB.NamedExec("delete from stacks where id = ?;", id)

// 	if err != nil {
// 		fmt.Printf("Failed deleting stack with id: %s because %s\n", id, err)
// 		return err
// 	}

// 	return nil
// }

func (stack *Stack) Remove() error {
	_, err := db.DB.NamedExec("delete from stacks where id = :id;", stack)

	if err != nil {
		fmt.Printf("Failed deleting stack with id: %s because %s\n", stack.ID, err)
		return err
	}

	return nil
}
