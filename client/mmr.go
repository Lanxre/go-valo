package client

import (
	"github.com/lanxre/go-valo/internal"
	"github.com/lanxre/go-valo/types/henrik_responses/mmr"
)

func (h *httpClient) GetMMRV2(options ...map[string]string) (mmr.MMRV2, error) {
	data, err := h.GetRaw("/v2/mmr/{region}/{name}/{tag}", options...)

	if err != nil {
		return mmr.MMRV2{}, err
	}

	return internal.ParseJSON[mmr.MMRV2](data)
}

func (h *httpClient) GetMMRByPUUIDV2(options ...map[string]string) (mmr.MMRV2, error) {
	data, err := h.GetRaw("/v2/by-puuid/mmr/{region}/{puuid}", options...)

	if err != nil {
		return mmr.MMRV2{}, err
	}

	return internal.ParseJSON[mmr.MMRV2](data)
}

func (h *httpClient) GetMMRV3(options ...map[string]string) (mmr.MMRV3, error) {
	data, err := h.GetRaw("/v3/mmr/{region}/{platform}/{name}/{tag}", options...)

	if err != nil {
		return mmr.MMRV3{}, err
	}

	return internal.ParseJSON[mmr.MMRV3](data)
}

func (h *httpClient) GetMMRByPUUIDV3(options ...map[string]string) (mmr.MMRV3, error) {
	data, err := h.GetRaw("/v3/by-puuid/mmr/{region}/{platform}/{puuid}", options...)

	if err != nil {
		return mmr.MMRV3{}, err
	}

	return internal.ParseJSON[mmr.MMRV3](data)
}
