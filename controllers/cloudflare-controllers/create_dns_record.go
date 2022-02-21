package cloudflareControllers

import (
	"fmt"
	"io"
	"net/http"
	"whm-api/utils/cloudflare"
	cloudflareDns "whm-api/utils/cloudflare/dns"
)

func (c controller) CreateDNSRecordController(zoneID string, body io.ReadCloser) (cloudflareDns.Record, string) {
	var record cloudflareDns.RecordCreatedResponse

	if err := cloudflare.Post("/zones/"+zoneID+"/dns_records", body, &record); err != nil {
		fmt.Println(err)

		return cloudflareDns.Record{}, "Couldn't fetch zone dns records."
	}

	return record.Result, http.StatusText(http.StatusOK)
}
