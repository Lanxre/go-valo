package leaderboard

import "time"

type Leaderboard struct {
	UpdatedAt  time.Time `json:"updated_at"`
	Thresholds []struct {
		Tier struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"tier"`
		StartIndex int `json:"start_index"`
		Threshold  int `json:"threshold"`
	} `json:"thresholds"`
	Players []struct {
		Card            string    `json:"card"`
		Title           string    `json:"title"`
		IsBanned        bool      `json:"is_banned"`
		IsAnonymized    bool      `json:"is_anonymized"`
		Puuid           string    `json:"puuid"`
		Name            string    `json:"name"`
		Tag             string    `json:"tag"`
		LeaderboardRank int       `json:"leaderboard_rank"`
		Tier            int       `json:"tier"`
		Rr              int       `json:"rr"`
		Wins            int       `json:"wins"`
		UpdatedAt       time.Time `json:"updated_at"`
	} `json:"players"`
}
