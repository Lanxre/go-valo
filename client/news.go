package client

import (
	"fmt"

	"github.com/lanxre/go-valo/internal"
	"github.com/lanxre/go-valo/types"
	"github.com/lanxre/go-valo/types/henrik_responses/website"
)

func (h *httpClient) GetWebsiteNews(params types.WebsiteNewsParams, options ...map[string]string) (website.WebsiteNews, error) {

	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	data, err := h.GetRaw(fmt.Sprintf("/v1/website/%s", params.Countrycode), options...)

	if err != nil {
		return website.WebsiteNews{}, err
	}

	return internal.ParseJSON[website.WebsiteNews](data)
}
