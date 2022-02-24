package stacks

import (
	"log"
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
		log.Printf("Failed deleting stack with id: %s because %s\n", stack.ID, err)
		return err
	}

	// dirName := stack.DirectoryName()

	// if err := os.RemoveAll("/data/stacks/" + dirName); err != nil {
	// 	log.Printf("Failed removing directory %s because: %s\n", dirName, err)
	// 	return err
	// }

	return nil
}
