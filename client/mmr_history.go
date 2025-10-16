package client

import (
	"github.com/lanxre/go-valo/internal"
	"github.com/lanxre/go-valo/types/henrik_responses/mmr_history"
)

func (h *httpClient) GetMMRHistoryV1(options ...map[string]string) (mmrhistory.MMRHistoryV1, error) {
	data, err := h.GetRaw("/v1/mmr-history/{region}/{name}/{tag}", options...)

	if err != nil {
		return mmrhistory.MMRHistoryV1{}, err
	}

	return internal.ParseJSON[mmrhistory.MMRHistoryV1](data)
}

func (h *httpClient) GetMMRHistoryByPUUIDV1(options ...map[string]string) (mmrhistory.MMRHistoryV1, error) {
	data, err := h.GetRaw("/v1/by-puuid/mmr-history/{region}/{puuid}", options...)

	if err != nil {
		return mmrhistory.MMRHistoryV1{}, err
	}

	return internal.ParseJSON[mmrhistory.MMRHistoryV1](data)
}

func (h *httpClient) GetMMRHistoryV2(options ...map[string]string) (mmrhistory.MMRHistoryV2, error) {
	data, err := h.GetRaw("/v2/mmr-history/{region}/{platform}/{name}/{tag}", options...)

	if err != nil {
		return mmrhistory.MMRHistoryV2{}, err
	}

	return internal.ParseJSON[mmrhistory.MMRHistoryV2](data)
}

func (h *httpClient) GetMMRHistoryByPUUIDV2(options ...map[string]string) (mmrhistory.MMRHistoryV2, error) {
	data, err := h.GetRaw("/v2/by-puuid/mmr-history/{region}/{platform}/{puuid}", options...)

	if err != nil {
		return mmrhistory.MMRHistoryV2{}, err
	}

	return internal.ParseJSON[mmrhistory.MMRHistoryV2](data)
}
