package client

import (
	"github.com/lanxre/go-valo/internal"
	"github.com/lanxre/go-valo/types/henrik_responses/leaderboard"
)

func (h *httpClient) GetLeaderboard(options ...map[string]string) (leaderboard.Leaderboard, error) {
	data, err := h.GetRaw("/v3/leaderboard/{region}/{platform}", options...)

	if err != nil {
		return leaderboard.Leaderboard{}, err
	}

	return internal.ParseJSON[leaderboard.Leaderboard](data)
}
