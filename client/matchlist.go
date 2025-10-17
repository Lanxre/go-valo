package client

import (
	"fmt"

	"github.com/lanxre/go-valo/internal"
	"github.com/lanxre/go-valo/types"
	"github.com/lanxre/go-valo/types/henrik_responses/matchlist"
)

func (h *httpClient) GetMatchesV3(params types.MatchlistParamV3, options ...map[string]string) (matchlist.MatchesV3, error) {

	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	data, err := h.GetRaw(fmt.Sprintf("/v3/matches/%s/%s/%s", params.Region, params.Name, params.Tag), options...)

	if err != nil {
		return matchlist.MatchesV3{}, err
	}

	return internal.ParseJSON[matchlist.MatchesV3](data)
}

func (h *httpClient) GetMatchesV4(params types.MatchlistParamV4, options ...map[string]string) (matchlist.MatchesV4, error) {

	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	data, err := h.GetRaw(fmt.Sprintf("/v4/matches/%s/%s/%s/%s", params.Region, params.Platform, params.Name, params.Tag), options...)

	if err != nil {
		return matchlist.MatchesV4{}, err
	}

	return internal.ParseJSON[matchlist.MatchesV4](data)
}

func (h *httpClient) GetMatchesByPUUIDV3(params types.PlayerRegionParams, options ...map[string]string) (matchlist.MatchesV3, error) {

	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	data, err := h.GetRaw(fmt.Sprintf("/v3/by-puuid/matches/%s/%s", params.Region, params.PUUID), options...)

	if err != nil {
		return matchlist.MatchesV3{}, err
	}

	return internal.ParseJSON[matchlist.MatchesV3](data)
}

func (h *httpClient) GetMatchesByPUUIDV4(params types.PlayerRegionPlatformParams, options ...map[string]string) (matchlist.MatchesV4, error) {

	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	data, err := h.GetRaw(fmt.Sprintf("/v4/by-puuid/matches/%s/%s/%s", params.Region, params.Platform, params.PUUID), options...)

	if err != nil {
		return matchlist.MatchesV4{}, err
	}

	return internal.ParseJSON[matchlist.MatchesV4](data)
}
