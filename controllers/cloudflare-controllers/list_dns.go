package cloudflareControllers

import (
	"fmt"
	"net/http"
	"whm-api/utils/cloudflare"
	cloudflareDns "whm-api/utils/cloudflare/dns"
)

func (c controller) ListDNSController(zoneID string) ([]cloudflareDns.Record, string) {
	var dns cloudflareDns.DNSResponse

	if err := cloudflare.Get("/zones/"+zoneID+"/dns_records", &dns); err != nil {
		fmt.Println(err)

		return nil, "Couldn't fetch zone dns records."
	}

	return dns.Result, http.StatusText(http.StatusOK)
}
