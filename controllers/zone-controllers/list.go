package zoneControllers

import (
	"fmt"
	"net/http"
	"whm-api/utils/db/zones"
)

func (c controller) List() ([]zones.Zone, string) {
	zonesList, err := zones.List()

	if err != nil {
		fmt.Println(err)
		return nil, "Couldn't get zones"
	}

	return zonesList, http.StatusText(http.StatusOK)
}
