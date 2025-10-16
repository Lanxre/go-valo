package client

import (
	"github.com/lanxre/go-valo/internal"
	"github.com/lanxre/go-valo/types/henrik_responses/esports"
)

func (h *httpClient) GetScheduleEsports(options ...map[string]string) (esports.ScheduleEsports, error) {
	data, err := h.GetRaw("/v1/esports/schedule", options...)

	if err != nil {
		return esports.ScheduleEsports{}, err
	}

	return internal.ParseJSON[esports.ScheduleEsports](data)
}
