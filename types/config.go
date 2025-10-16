package types

import (
	"net/http"
	"time"
)

type Config struct {
	HenrikBaseURL string
	HenrikAPIKey  string
	Timeout       time.Duration
	HTTPClient    *http.Client
}
