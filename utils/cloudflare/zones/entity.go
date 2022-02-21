package cloudflareZones

import "time"

type ZonesResponse struct {
	Result     []Zone `json:"result"`
	ResultInfo struct {
		Page       int `json:"page"`
		PerPage    int `json:"per_page"`
		TotalPages int `json:"total_pages"`
		Count      int `json:"count"`
		TotalCount int `json:"total_count"`
	} `json:"result_info"`
	Success  bool          `json:"success"`
	Errors   []interface{} `json:"errors"`
	Messages []interface{} `json:"messages"`
}

type Zone struct {
	ID                  string      `json:"id"`
	Name                string      `json:"name"`
	Status              string      `json:"status"`
	Paused              bool        `json:"paused"`
	Type                string      `json:"type"`
	DevelopmentMode     int         `json:"development_mode"`
	NameServers         []string    `json:"name_servers"`
	OriginalNameServers []string    `json:"original_name_servers"`
	OriginalRegistrar   interface{} `json:"original_registrar"`
	OriginalDnshost     interface{} `json:"original_dnshost"`
	ModifiedOn          time.Time   `json:"modified_on"`
	CreatedOn           time.Time   `json:"created_on"`
	ActivatedOn         time.Time   `json:"activated_on"`
	Meta                struct {
		Step                    int  `json:"step"`
		WildcardProxiable       bool `json:"wildcard_proxiable"`
		CustomCertificateQuota  int  `json:"custom_certificate_quota"`
		PageRuleQuota           int  `json:"page_rule_quota"`
		PhishingDetected        bool `json:"phishing_detected"`
		MultipleRailgunsAllowed bool `json:"multiple_railguns_allowed"`
	} `json:"meta"`
	Owner struct {
		ID    string `json:"id"`
		Type  string `json:"type"`
		Email string `json:"email"`
	} `json:"owner"`
	Account struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"account"`
	Permissions []string `json:"permissions"`
	Plan        struct {
		ID                string `json:"id"`
		Name              string `json:"name"`
		Price             int    `json:"price"`
		Currency          string `json:"currency"`
		Frequency         string `json:"frequency"`
		IsSubscribed      bool   `json:"is_subscribed"`
		CanSubscribe      bool   `json:"can_subscribe"`
		LegacyID          string `json:"legacy_id"`
		LegacyDiscount    bool   `json:"legacy_discount"`
		ExternallyManaged bool   `json:"externally_managed"`
	} `json:"plan"`
}
