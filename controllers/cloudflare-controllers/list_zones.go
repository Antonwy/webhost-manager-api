package cloudflareControllers

import (
	"fmt"
	"net/http"
	"whm-api/utils/cloudflare"
	cloudflareZones "whm-api/utils/cloudflare/zones"
)

func (c controller) ListZonesController() ([]cloudflareZones.Zone, string) {
	var zones cloudflareZones.ZonesResponse

	if err := cloudflare.Get("/zones", &zones); err != nil {
		fmt.Println(err)

		return nil, "Couldn't fetch cloudflare zones."
	}

	return zones.Result, http.StatusText(http.StatusOK)
}
