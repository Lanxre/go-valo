package client

import (
	"github.com/lanxre/go-valo/internal"
	"github.com/lanxre/go-valo/types/henrik_responses/match"
)

func (h *httpClient) GetMatchV2(options ...map[string]string) (match.MatchV2, error) {
	data, err := h.GetRaw("/v2/match/{matchid}", options...)

	if err != nil {
		return match.MatchV2{}, err
	}

	return internal.ParseJSON[match.MatchV2](data)
}

func (h *httpClient) GetMatchV4(options ...map[string]string) (match.MatchV4, error) {
	data, err := h.GetRaw("/v4/match/{region}/{matchid}", options...)

	if err != nil {
		return match.MatchV4{}, err
	}

	return internal.ParseJSON[match.MatchV4](data)
}
