package client

import (
	"fmt"

	"github.com/lanxre/go-valo/internal"
	"github.com/lanxre/go-valo/types"
	"github.com/lanxre/go-valo/types/henrik_responses/mmr_history"
)

func (h *httpClient) GetMMRHistoryV1(params types.MatchlistParamV3, options ...map[string]string) (mmrhistory.MMRHistoryV1, error) {

	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	data, err := h.GetRaw(fmt.Sprintf("/v1/mmr-history/%s/%s/%s", params.Region, params.Name, params.Tag), options...)

	if err != nil {
		return mmrhistory.MMRHistoryV1{}, err
	}

	return internal.ParseJSON[mmrhistory.MMRHistoryV1](data)
}

func (h *httpClient) GetMMRHistoryByPUUIDV1(params types.PlayerRegionParams, options ...map[string]string) (mmrhistory.MMRHistoryV1, error) {

	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	data, err := h.GetRaw(fmt.Sprintf("/v1/by-puuid/mmr-history/%s/%s", params.Region, params.PUUID), options...)

	if err != nil {
		return mmrhistory.MMRHistoryV1{}, err
	}

	return internal.ParseJSON[mmrhistory.MMRHistoryV1](data)
}

func (h *httpClient) GetMMRHistoryV2(params types.MatchlistParamV4, options ...map[string]string) (mmrhistory.MMRHistoryV2, error) {

	if err := params.Validate(); err != nil {
		return mmrhistory.MMRHistoryV2{}, fmt.Errorf("validation failed: %w", err)
	}

	data, err := h.GetRaw(fmt.Sprintf("/v2/mmr-history/%s/%s/%s/%s", params.Region, params.Platform, params.Name, params.Tag), options...)

	if err != nil {
		return mmrhistory.MMRHistoryV2{}, err
	}

	return internal.ParseJSON[mmrhistory.MMRHistoryV2](data)
}

func (h *httpClient) GetMMRHistoryByPUUIDV2(params types.PlayerRegionPlatformParams, options ...map[string]string) (mmrhistory.MMRHistoryV2, error) {

	if err := params.Validate(); err != nil {
		return mmrhistory.MMRHistoryV2{}, fmt.Errorf("validation failed: %w", err)
	}

	data, err := h.GetRaw(fmt.Sprintf("/v2/by-puuid/mmr-history/%s/%s/%s", params.Region, params.Platform, params.PUUID), options...)

	if err != nil {
		return mmrhistory.MMRHistoryV2{}, err
	}

	return internal.ParseJSON[mmrhistory.MMRHistoryV2](data)
}
