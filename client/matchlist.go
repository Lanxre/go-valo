package client

import (
	"github.com/lanxre/go-valo/internal"
	"github.com/lanxre/go-valo/types/henrik_responses/matchlist"
)

func (h *httpClient) GetMatchesV3(options ...map[string]string) (matchlist.MatchesV3, error) {
	data, err := h.GetRaw("/v3/matches/{region}/{name}/{tag}", options...)

	if err != nil {
		return matchlist.MatchesV3{}, err
	}

	return internal.ParseJSON[matchlist.MatchesV3](data)
}

func (h *httpClient) GetMatchesV4(options ...map[string]string) (matchlist.MatchesV4, error) {
	data, err := h.GetRaw("/v4/matches/{region}/{platform}/{name}/{tag}", options...)

	if err != nil {
		return matchlist.MatchesV4{}, err
	}

	return internal.ParseJSON[matchlist.MatchesV4](data)
}

func (h *httpClient) GetMatchesByPUUIDV3(options ...map[string]string) (matchlist.MatchesV3, error) {
	data, err := h.GetRaw("/v3/by-puuid/matches/{region}/{puuid}", options...)

	if err != nil {
		return matchlist.MatchesV3{}, err
	}

	return internal.ParseJSON[matchlist.MatchesV3](data)
}

func (h *httpClient) GetMatchesByPUUIDV4(options ...map[string]string) (matchlist.MatchesV4, error) {
	data, err := h.GetRaw("/v4/by-puuid/matches/{region}/{platform}/{puuid}", options...)

	if err != nil {
		return matchlist.MatchesV4{}, err
	}

	return internal.ParseJSON[matchlist.MatchesV4](data)
}
