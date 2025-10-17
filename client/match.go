package client

import (
	"fmt"

	"github.com/lanxre/go-valo/internal"
	"github.com/lanxre/go-valo/types"
	"github.com/lanxre/go-valo/types/henrik_responses/match"
)

func (h *httpClient) GetMatchV2(params types.MatchParamV2, options ...map[string]string) (match.MatchV2, error) {

	if err := params.Validate(); err != nil {
		return match.MatchV2{}, fmt.Errorf("validation failed: %w", err)
	}

	data, err := h.GetRaw(fmt.Sprintf("/v2/match/%s", params.Uuid), options...)

	if err != nil {
		return match.MatchV2{}, err
	}

	return internal.ParseJSON[match.MatchV2](data)
}

func (h *httpClient) GetMatchV4(params types.MatchParamV4, options ...map[string]string) (match.MatchV4, error) {

	if err := params.Validate(); err != nil {
		return match.MatchV4{}, fmt.Errorf("validation failed: %w", err)
	}

	data, err := h.GetRaw(fmt.Sprintf("/v4/match/%s/%s", params.Region, params.Uuid), options...)

	if err != nil {
		return match.MatchV4{}, err
	}

	return internal.ParseJSON[match.MatchV4](data)
}
