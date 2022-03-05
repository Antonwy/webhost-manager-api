package zoneControllers

import (
	"fmt"
	"net/http"
	"whm-api/utils/db/zones"
)

func (c controller) Remove(id string) string {
	err := zones.RemoveFromId(id)

	if err != nil {
		fmt.Println(err)
		return "Couldn't remove zone with id: " + id
	}

	return http.StatusText(http.StatusOK)
}
