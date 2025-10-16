package client

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/lanxre/go-valo/errors"
	"github.com/lanxre/go-valo/internal"
	"github.com/lanxre/go-valo/types/henrik_responses/accounts"
)



type httpClient struct {
	henrikApiBaseURL string
	henrikApiKey     string
	httpClient       *http.Client
}

func (h *httpClient) GetAccountV1(options ...map[string]string) (accounts.AccountV1, error) {
    data, err := h.GetRaw("/v1/account/{name}/{tag}", options...)
    
	if err != nil {
        return accounts.AccountV1{}, err
    }
    
    return internal.ParseJSON[accounts.AccountV1](data)
}

func (h *httpClient) GetRaw(path string, options ...map[string]string) ([]byte, error) {
    return h.GetRawWithContext(context.Background(), path, options...)
}


func (h *httpClient) GetRawWithContext(ctx context.Context, path string, options ...map[string]string) ([]byte, error) {
    url := h.henrikApiBaseURL + path

    if len(options) > 0 && len(options[0]) > 0 {
        pathParams := options[0]
        url = internal.ReplaceNamedPathParams(url, pathParams)
    }

	var queryParams map[string]string
    if len(options) > 1 && len(options[1]) > 0 {
        queryParams = options[1]
    }

	if len(queryParams) > 0 {
        url += "?" + internal.BuildQueryString(queryParams)
    }


    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to create request: %w", err)
    }

    req.Header.Set("Authorization", h.henrikApiKey)
    req.Header.Set("accept", "application/json")
	
    resp, err := h.httpClient.Do(req)
    if err != nil {
        return nil, fmt.Errorf("request failed: %w", err)
    }
    
	defer resp.Body.Close()

    if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
        body, _ := io.ReadAll(resp.Body)
        return nil, &errors.APIError{
            StatusCode: resp.StatusCode,
            Message:    resp.Status,
            Body:       string(body),
        }
    }

    data, err := io.ReadAll(resp.Body)

    if err != nil {
        return nil, fmt.Errorf("failed to read response body: %w", err)
    }

    return data, nil
}