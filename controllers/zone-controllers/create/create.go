package createZoneController

import (
	"log"
	"net/http"
	"whm-api/utils/db/zones"
)

func (c controller) Create(input InputCreateZone) (zones.Zone, string) {

	zone := zones.Zone{
		Name:                 input.Name,
		ID:                   input.ID,
		SyncedWithCloudflare: true,
	}

	if err := zone.Create(); err != nil {
		log.Println(err)
		return zones.Zone{}, "Couldn't create zone!"
	}

	return zone, http.StatusText(http.StatusOK)
}
