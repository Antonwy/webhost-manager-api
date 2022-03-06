package cloudflareControllers

import (
	"io"
	cloudflareDns "whm-api/utils/cloudflare/dns"
	cloudflareZones "whm-api/utils/cloudflare/zones"
)

type Controller interface {
	ListZonesController() ([]cloudflareZones.Zone, string)
	ListDNSController(zoneID string) ([]cloudflareDns.Record, string)
	CreateDNSRecordController(zoneID string, body io.ReadCloser) (cloudflareDns.Record, string)
	DeleteRecordController(zoneID string, recordID string) string
}

type controller struct{}

func NewController() *controller {
	return &controller{}
}
