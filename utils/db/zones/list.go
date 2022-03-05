package zones

import "whm-api/utils/db"

func List() ([]Zone, error) {
	var zones []Zone

	if err := db.DB.Select(&zones, "select * from zones"); err != nil {
		return nil, err
	}

	return zones, nil
}
