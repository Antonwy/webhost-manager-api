package stacks

import "whm-api/utils/db"

func GetFromID(id string) (Stack, error) {
	stack := Stack{}
	if err := db.DB.Get(&stack, "select * from stacks where id = $1;", id); err != nil {
		return Stack{}, err
	}

	return stack, nil
}
