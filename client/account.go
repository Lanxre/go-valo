package client

import (
	"fmt"

	"github.com/lanxre/go-valo/internal"
	"github.com/lanxre/go-valo/types"
	"github.com/lanxre/go-valo/types/henrik_responses/accounts"
)

func (h *httpClient) GetAccountV1(params types.ValorantAccountParams, options ...map[string]string) (accounts.AccountV1, error) {

	if err := params.Validate(); err != nil {
		return accounts.AccountV1{}, fmt.Errorf("validation failed: %w", err)
	}

	data, err := h.GetRaw(fmt.Sprintf("/v1/account/%s/%s", params.Name, params.Tag), options...)

	if err != nil {
		return accounts.AccountV1{}, err
	}

	return internal.ParseJSON[accounts.AccountV1](data)
}

func (h *httpClient) GetAccountV2(params types.ValorantAccountParams, options ...map[string]string) (accounts.AccountV2, error) {

	if err := params.Validate(); err != nil {
		return accounts.AccountV2{}, fmt.Errorf("validation failed: %w", err)
	}

	data, err := h.GetRaw(fmt.Sprintf("/v2/account/%s/%s", params.Name, params.Tag), options...)

	if err != nil {
		return accounts.AccountV2{}, err
	}

	return internal.ParseJSON[accounts.AccountV2](data)
}

func (h *httpClient) GetAccountByPUUIDV1(params types.PlayerParams, options ...map[string]string) (accounts.AccountV1, error) {

	if err := params.Validate(); err != nil {
		return accounts.AccountV1{}, fmt.Errorf("validation failed: %w", err)
	}

	data, err := h.GetRaw(fmt.Sprintf("/v1/by-puuid/account/%s", params.PUUID), options...)

	if err != nil {
		return accounts.AccountV1{}, err
	}

	return internal.ParseJSON[accounts.AccountV1](data)
}

func (h *httpClient) GetAccountByPUUIDV2(params types.PlayerParams, options ...map[string]string) (accounts.AccountV1, error) {

	if err := params.Validate(); err != nil {
		return accounts.AccountV1{}, fmt.Errorf("validation failed: %w", err)
	}

	data, err := h.GetRaw(fmt.Sprintf("/v2/by-puuid/account/%s", params.PUUID), options...)

	if err != nil {
		return accounts.AccountV1{}, err
	}

	return internal.ParseJSON[accounts.AccountV1](data)
}
