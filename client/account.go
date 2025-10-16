package client

import (
	"github.com/lanxre/go-valo/internal"
	"github.com/lanxre/go-valo/types/henrik_responses/accounts"
)

func (h *httpClient) GetAccountV1(options ...map[string]string) (accounts.AccountV1, error) {
    data, err := h.GetRaw("/v1/account/{name}/{tag}", options...)
    
	if err != nil {
        return accounts.AccountV1{}, err
    }
    
    return internal.ParseJSON[accounts.AccountV1](data)
}

func (h *httpClient) GetAccountV2(options ...map[string]string) (accounts.AccountV2, error) {
    data, err := h.GetRaw("/v2/account/{name}/{tag}", options...)
    
	if err != nil {
        return accounts.AccountV2{}, err
    }
    
    return internal.ParseJSON[accounts.AccountV2](data)
}

func (h *httpClient) GetAccountByPUUIDV1(options ...map[string]string) (accounts.AccountV1, error) {
    data, err := h.GetRaw("/v1/by-puuid/account/{puuid}", options...)
    
	if err != nil {
        return accounts.AccountV1{}, err
    }
    
    return internal.ParseJSON[accounts.AccountV1](data)
}

func (h *httpClient) GetAccountByPUUIDV2(options ...map[string]string) (accounts.AccountV1, error) {
    data, err := h.GetRaw("/v2/by-puuid/account/{puuid}", options...)
    
	if err != nil {
        return accounts.AccountV1{}, err
    }
    
    return internal.ParseJSON[accounts.AccountV1](data)
}