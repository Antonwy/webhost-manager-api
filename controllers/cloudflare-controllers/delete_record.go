package cloudflareControllers

import (
	"fmt"
	"net/http"
	"whm-api/utils/cloudflare"
	cloudflareDns "whm-api/utils/cloudflare/dns"
)

func (c controller) DeleteRecordController(zoneID string, recordID string) string {
	var res cloudflareDns.RecordDeletedResponse
	if err := cloudflare.Delete("/zones/"+zoneID+"/dns_records/"+recordID, &res); err != nil {
		fmt.Println(err)
		return "Couldn't fetch zone dns records."
	}

	return http.StatusText(http.StatusOK)
}
