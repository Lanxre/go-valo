package website

import "time"

type WebsiteNews []struct {
	BannerURL    string    `json:"banner_url"`
	Category     string    `json:"category"`
	Date         time.Time `json:"date"`
	ExternalLink string    `json:"external_link"`
	Title        string    `json:"title"`
	URL          string    `json:"url"`
}
