package client

import (
	"github.com/lanxre/go-valo/internal"
	"github.com/lanxre/go-valo/types/henrik_responses/content"
)

func (h *httpClient) GetContent(options ...map[string]string) (content.Content, error) {
	data, err := h.GetRaw("/v1/content", options...)

	if err != nil {
		return content.Content{}, err
	}

	return internal.ParseJSON[content.Content](data)
}
