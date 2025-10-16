package client

import (
	"github.com/lanxre/go-valo/internal"
	"github.com/lanxre/go-valo/types/henrik_responses/website"
)

func (h *httpClient) GetWebsiteNews(options ...map[string]string) (website.WebsiteNews, error) {
	data, err := h.GetRaw("/v1/website/{countrycode}", options...)

	if err != nil {
		return website.WebsiteNews{}, err
	}

	return internal.ParseJSON[website.WebsiteNews](data)
}
