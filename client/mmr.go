package client

import (
	"fmt"

	"github.com/lanxre/go-valo/internal"
	"github.com/lanxre/go-valo/types"
	"github.com/lanxre/go-valo/types/henrik_responses/mmr"
)

func (h *httpClient) GetMMRV2(params types.MatchlistParamV3, options ...map[string]string) (mmr.MMRV2, error) {

	if err := params.Validate(); err != nil {
		return mmr.MMRV2{}, fmt.Errorf("validation failed: %w", err)
	}

	data, err := h.GetRaw(fmt.Sprintf("/v2/mmr/%s/%s/%s", params.Region, params.Name, params.Tag), options...)

	if err != nil {
		return mmr.MMRV2{}, err
	}

	return internal.ParseJSON[mmr.MMRV2](data)
}

func (h *httpClient) GetMMRByPUUIDV2(params types.PlayerRegionParams, options ...map[string]string) (mmr.MMRV2, error) {

	if err := params.Validate(); err != nil {
		return mmr.MMRV2{}, fmt.Errorf("validation failed: %w", err)
	}

	data, err := h.GetRaw(fmt.Sprintf("/v2/by-puuid/mmr/%s/%s", params.Region, params.PUUID), options...)

	if err != nil {
		return mmr.MMRV2{}, err
	}

	return internal.ParseJSON[mmr.MMRV2](data)
}

func (h *httpClient) GetMMRV3(params types.MatchlistParamV4, options ...map[string]string) (mmr.MMRV3, error) {

	if err := params.Validate(); err != nil {
		return mmr.MMRV3{}, fmt.Errorf("validation failed: %w", err)
	}

	data, err := h.GetRaw(fmt.Sprintf("/v3/mmr/%s/%s/%s/%s", params.Region, params.Platform, params.Name, params.Tag), options...)

	if err != nil {
		return mmr.MMRV3{}, err
	}

	return internal.ParseJSON[mmr.MMRV3](data)
}

func (h *httpClient) GetMMRByPUUIDV3(params types.PlayerRegionPlatformParams, options ...map[string]string) (mmr.MMRV3, error) {

	if err := params.Validate(); err != nil {
		return mmr.MMRV3{}, fmt.Errorf("validation failed: %w", err)
	}

	data, err := h.GetRaw(fmt.Sprintf("/v3/by-puuid/mmr/%s/%s/%s", params.Region, params.Platform, params.PUUID), options...)

	if err != nil {
		return mmr.MMRV3{}, err
	}

	return internal.ParseJSON[mmr.MMRV3](data)
}
