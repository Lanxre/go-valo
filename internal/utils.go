package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/lanxre/go-valo/types"
)

func ParseJSON[T any](data []byte) (T, error) {
    var resp types.Response[T]
    if err := json.Unmarshal(data, &resp); err != nil {
        var zero T
        return zero, fmt.Errorf("JSON parse error: %w", err)
    }
    
    if resp.Status != http.StatusOK {
        var zero T
        return zero, fmt.Errorf("API error: %s", resp.Error)
    }
    
    return resp.Data, nil
}

func ReplacePathParams(path string, params []string) string {
    result := path
    for _, param := range params {
        result = strings.Replace(result, "{}", param, 1)
    }
    return result
}

func ReplaceNamedPathParams(path string, params map[string]string) string {
    result := path
    for key, value := range params {
        placeholder := "{" + key + "}"
        result = strings.Replace(result, placeholder, value, 1)
    }
    return result
}

func BuildQueryString(params map[string]string) string {
    if len(params) == 0 {
        return ""
    }
    
    var queryParts []string
    for key, value := range params {
        encodedKey := url.QueryEscape(key)
        encodedValue := url.QueryEscape(value)
        queryParts = append(queryParts, encodedKey+"="+encodedValue)
    }
    return strings.Join(queryParts, "&")
}