package zones

import "whm-api/utils/db"

func GetFromID(id string) (Zone, error) {
	zone := Zone{}
	if err := db.DB.Get(&zone, "select * from zones where id = $1", id); err != nil {
		return Zone{}, err
	}

	return zone, nil
}

func GetFromName(name string) (Zone, error) {
	zone := Zone{}
	if err := db.DB.Get(&zone, "select * from zones where name = $1", name); err != nil {
		return Zone{}, err
	}

	return zone, nil
}
