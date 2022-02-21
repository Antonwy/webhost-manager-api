package cloudflare

import (
	"encoding/json"
	"io"
	"net/http"
	util "whm-api/utils"
)

func Get(route string, decode interface{}) error {
	req, err := http.NewRequest(http.MethodGet, URL+route, nil)

	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Bearer "+util.GodotEnv("CLOUDFLARE_TOKEN"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if err := json.NewDecoder(res.Body).Decode(&decode); err != nil {
		return err
	}

	return nil
}

func Post(route string, body io.Reader, decode interface{}) error {
	req, err := http.NewRequest(http.MethodPost, URL+route, body)

	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Bearer "+util.GodotEnv("CLOUDFLARE_TOKEN"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if err := json.NewDecoder(res.Body).Decode(&decode); err != nil {
		return err
	}

	return nil
}
