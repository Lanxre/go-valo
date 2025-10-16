package client

import (
	"github.com/lanxre/go-valo/internal"
	"github.com/lanxre/go-valo/types/henrik_responses/stored_matches"
)

func (h *httpClient) GetStoredMMRHistoryByPUUIDV1(options ...map[string]string) (storedmatches.StoredMMRHistoryV1, error) {
	data, err := h.GetRaw("/v1/by-puuid/stored-mmr-history/{region}/{puuid}", options...)

	if err != nil {
		return storedmatches.StoredMMRHistoryV1{}, err
	}

	return internal.ParseJSON[storedmatches.StoredMMRHistoryV1](data)
}

func (h *httpClient) GetStoredMMRHistoryByPUUIDV2(options ...map[string]string) (storedmatches.StoredMMRHistoryV2, error) {
	data, err := h.GetRaw("/v2/by-puuid/stored-mmr-history/{region}/{platform}/{puuid}", options...)

	if err != nil {
		return storedmatches.StoredMMRHistoryV2{}, err
	}

	return internal.ParseJSON[storedmatches.StoredMMRHistoryV2](data)
}

func (h *httpClient) GetStoredMMRHistoryV1(options ...map[string]string) (storedmatches.StoredMMRHistoryV1, error) {
	data, err := h.GetRaw("/v1/stored-mmr-history/{region}/{name}/{tag}", options...)

	if err != nil {
		return storedmatches.StoredMMRHistoryV1{}, err
	}

	return internal.ParseJSON[storedmatches.StoredMMRHistoryV1](data)
}

func (h *httpClient) GetStoredMMRHistoryV2(options ...map[string]string) (storedmatches.StoredMMRHistoryV2, error) {
	data, err := h.GetRaw("/v2/stored-mmr-history/{region}/{platform}/{name}/{tag}", options...)

	if err != nil {
		return storedmatches.StoredMMRHistoryV2{}, err
	}

	return internal.ParseJSON[storedmatches.StoredMMRHistoryV2](data)
}

func (h *httpClient) GetStoredMatchesPUUIDV1(options ...map[string]string) (storedmatches.StoredMatchesV1, error) {
	data, err := h.GetRaw("/v1/by-puuid/stored-matches/{region}/{puuid}", options...)

	if err != nil {
		return storedmatches.StoredMatchesV1{}, err
	}

	return internal.ParseJSON[storedmatches.StoredMatchesV1](data)
}

func (h *httpClient) GetStoredMatchesV1(options ...map[string]string) (storedmatches.StoredMatchesV1, error) {
	data, err := h.GetRaw("/v1/stored-matches/{region}/{name}/{tag}", options...)

	if err != nil {
		return storedmatches.StoredMatchesV1{}, err
	}

	return internal.ParseJSON[storedmatches.StoredMatchesV1](data)
}