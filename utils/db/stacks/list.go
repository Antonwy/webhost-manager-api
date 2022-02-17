package stacks

import "whm-api/utils/db"

func ListStacks() ([]Stack, error) {
	stacks := []Stack{}
	err := db.DB.Select(&stacks, "select * from stacks;")

	if err != nil {
		return nil, err
	}

	return nil, nil
}
