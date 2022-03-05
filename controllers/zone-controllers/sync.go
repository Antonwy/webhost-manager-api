package zoneControllers

import (
	"fmt"
	"log"
	"net/http"
	cloudflareControllers "whm-api/controllers/cloudflare-controllers"
	cloudflareZones "whm-api/utils/cloudflare/zones"
	"whm-api/utils/db/zones"
)

func (c controller) Sync() ([]zones.Zone, string) {
	cController := cloudflareControllers.NewController()

	cZones, err := cController.ListZonesController()

	if err != http.StatusText(http.StatusOK) {
		log.Println(err)
		return nil, "Couldn't fetch cloudflare zones"
	}

	dbZones, dbErr := zones.List()
	if dbErr != nil {
		fmt.Println(err)
		return nil, "Couldn't fetch db zones"
	}

	for _, zone := range dbZones {
		extractedCloudZone := extractZone(zone, cZones)
		if extractedCloudZone == nil {
			log.Printf("Couldn't find cloudflare zone with id: %s setting db zone to unsynced!\n", zone.ID)
			zone.SyncedWithCloudflare = false
			err := zone.Update()
			if err != nil {
				log.Println(err)
				continue
			}
		} else if zone.SyncedWithCloudflare == false {
			log.Printf("Found cloudflare zone with id: %s, but db zone was not synced.\n", extractedCloudZone.ID)
			zone.SyncedWithCloudflare = true
			err := zone.Update()
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}

	syncedZones, syncedDbErr := zones.List()
	if syncedDbErr != nil {
		fmt.Println(err)
		return nil, "Couldn't fetch synced db zones"
	}

	return syncedZones, http.StatusText(http.StatusOK)
}

func extractZone(zone zones.Zone, cloudZones []cloudflareZones.Zone) *cloudflareZones.Zone {
	for _, cloudZone := range cloudZones {
		if zone.ID == cloudZone.ID {
			return &cloudZone
		}
	}

	return nil
}
