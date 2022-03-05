package zoneControllers

import (
	"whm-api/utils/db/zones"
)

type Controller interface {
	Get(id string) (zones.Zone, string)
	List() ([]zones.Zone, string)
	Sync() ([]zones.Zone, string)
	Remove(id string) string
}

type controller struct{}

func NewController() *controller {
	return &controller{}
}
