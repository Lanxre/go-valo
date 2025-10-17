package client

import (
	"fmt"

	"github.com/lanxre/go-valo/internal"
	"github.com/lanxre/go-valo/types"
	"github.com/lanxre/go-valo/types/henrik_responses/stored_matches"
)

func (h *httpClient) GetStoredMMRHistoryByPUUIDV1(params types.PlayerRegionParams, options ...map[string]string) (storedmatches.StoredMMRHistoryV1, error) {

	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	data, err := h.GetRaw(fmt.Sprintf("/v1/by-puuid/stored-mmr-history/%s/%s", params.Region, params.PUUID), options...)

	if err != nil {
		return storedmatches.StoredMMRHistoryV1{}, err
	}

	return internal.ParseJSON[storedmatches.StoredMMRHistoryV1](data)
}

func (h *httpClient) GetStoredMMRHistoryByPUUIDV2(params types.PlayerRegionPlatformParams, options ...map[string]string) (storedmatches.StoredMMRHistoryV2, error) {

	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	data, err := h.GetRaw(fmt.Sprintf("/v2/by-puuid/stored-mmr-history/%s/%s/%s", params.Region, params.Platform, params.PUUID), options...)

	if err != nil {
		return storedmatches.StoredMMRHistoryV2{}, err
	}

	return internal.ParseJSON[storedmatches.StoredMMRHistoryV2](data)
}

func (h *httpClient) GetStoredMMRHistoryV1(params types.MatchlistParamV3, options ...map[string]string) (storedmatches.StoredMMRHistoryV1, error) {

	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	data, err := h.GetRaw(fmt.Sprintf("/v1/stored-mmr-history/%s/%s/%s", params.Region, params.Name, params.Tag), options...)

	if err != nil {
		return storedmatches.StoredMMRHistoryV1{}, err
	}

	return internal.ParseJSON[storedmatches.StoredMMRHistoryV1](data)
}

func (h *httpClient) GetStoredMMRHistoryV2(params types.MatchlistParamV4, options ...map[string]string) (storedmatches.StoredMMRHistoryV2, error) {

	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	data, err := h.GetRaw(fmt.Sprintf("/v2/stored-mmr-history/%s/%s/%s/%s", params.Region, params.Platform, params.Name, params.Tag), options...)

	if err != nil {
		return storedmatches.StoredMMRHistoryV2{}, err
	}

	return internal.ParseJSON[storedmatches.StoredMMRHistoryV2](data)
}

func (h *httpClient) GetStoredMatchesPUUIDV1(params types.PlayerRegionParams, options ...map[string]string) (storedmatches.StoredMatchesV1, error) {

	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	data, err := h.GetRaw(fmt.Sprintf("/v1/by-puuid/stored-matches/%s/%s", params.Region, params.PUUID), options...)

	if err != nil {
		return storedmatches.StoredMatchesV1{}, err
	}

	return internal.ParseJSON[storedmatches.StoredMatchesV1](data)
}

func (h *httpClient) GetStoredMatchesV1(params types.MatchlistParamV3, options ...map[string]string) (storedmatches.StoredMatchesV1, error) {

	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	data, err := h.GetRaw(fmt.Sprintf("/v1/stored-matches/%s/%s/%s", params.Region, params.Name, params.Tag), options...)

	if err != nil {
		return storedmatches.StoredMatchesV1{}, err
	}

	return internal.ParseJSON[storedmatches.StoredMatchesV1](data)
}
