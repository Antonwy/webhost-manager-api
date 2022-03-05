package zones

import (
	"fmt"
	"whm-api/utils/db"
)

func (zone Zone) Remove() error {
	_, err := db.DB.NamedExec("delete from zones where id = :id;", zone)

	if err != nil {
		fmt.Printf("Failed removing zone: %s because: %s\n", zone, err)
		return err
	}

	return nil
}

func RemoveFromId(id string) error {
	_, err := db.DB.Exec("delete from zones where id = $1;", id)

	if err != nil {
		fmt.Printf("Failed removing zone: %s because: %s\n", id, err)
		return err
	}

	return nil
}
