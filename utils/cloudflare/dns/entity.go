package cloudflareDns

import "time"

type DNSResponse struct {
	Result     []Record      `json:"result"`
	Success    bool          `json:"success"`
	Errors     []interface{} `json:"errors"`
	Messages   []interface{} `json:"messages"`
	ResultInfo struct {
		Page       int `json:"page"`
		PerPage    int `json:"per_page"`
		Count      int `json:"count"`
		TotalCount int `json:"total_count"`
		TotalPages int `json:"total_pages"`
	} `json:"result_info"`
}

type Record struct {
	ID        string `json:"id"`
	ZoneID    string `json:"zone_id"`
	ZoneName  string `json:"zone_name"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Content   string `json:"content"`
	Proxiable bool   `json:"proxiable"`
	Proxied   bool   `json:"proxied"`
	TTL       int    `json:"ttl"`
	Locked    bool   `json:"locked"`
	Meta      struct {
		AutoAdded           bool   `json:"auto_added"`
		ManagedByApps       bool   `json:"managed_by_apps"`
		ManagedByArgoTunnel bool   `json:"managed_by_argo_tunnel"`
		Source              string `json:"source"`
	} `json:"meta"`
	CreatedOn  time.Time `json:"created_on"`
	ModifiedOn time.Time `json:"modified_on"`
	Priority   int       `json:"priority,omitempty"`
}

type RecordCreatedResponse struct {
	Result   Record        `json:"result"`
	Success  bool          `json:"success"`
	Errors   []interface{} `json:"errors"`
	Messages []interface{} `json:"messages"`
}
