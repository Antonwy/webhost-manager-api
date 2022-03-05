package createZoneController

import (
	"whm-api/utils/db/zones"
)

type Controller interface {
	Create(zone InputCreateZone) (zones.Zone, string)
}

type controller struct{}

func NewController() *controller {
	return &controller{}
}
