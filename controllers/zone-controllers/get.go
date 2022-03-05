package zoneControllers

import (
	"fmt"
	"net/http"
	"whm-api/utils/db/zones"
)

func (c controller) Get(id string) (zones.Zone, string) {
	zone, err := zones.GetFromID(id)

	if err != nil {
		fmt.Println(err)
		return zones.Zone{}, "Couldn't get zone with id: " + id
	}

	return zone, http.StatusText(http.StatusOK)
}
