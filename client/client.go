package client

import (
	"net/http"
	"time"

	"github.com/lanxre/go-valo/types"
)

func NewHenrikClient(config types.Config) types.HenrikClient {
	config.HenrikBaseURL = "https://api.henrikdev.xyz/valorant"

	if config.HTTPClient == nil {
		config.HTTPClient = &http.Client{
			Timeout: config.Timeout,
		}
	}

	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second
	}

	return &httpClient{
		henrikApiBaseURL: config.HenrikBaseURL,
		henrikApiKey:     config.HenrikAPIKey,
		httpClient:       config.HTTPClient,
	}
}
