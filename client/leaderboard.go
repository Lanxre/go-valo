package client

import (
	"fmt"

	"github.com/lanxre/go-valo/internal"
	"github.com/lanxre/go-valo/types"
	"github.com/lanxre/go-valo/types/henrik_responses/leaderboard"
)

func (h *httpClient) GetLeaderboard(params types.LeaderboardParams, options ...map[string]string) (leaderboard.Leaderboard, error) {

	if err := params.Validate(); err != nil {
		return leaderboard.Leaderboard{}, fmt.Errorf("validation failed: %w", err)
	}

	data, err := h.GetRaw(fmt.Sprintf("/v3/leaderboard/%s/%s", params.Region, params.Platform), options...)

	if err != nil {
		return leaderboard.Leaderboard{}, err
	}

	return internal.ParseJSON[leaderboard.Leaderboard](data)
}
