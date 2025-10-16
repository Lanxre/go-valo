package mmr

import "time"

type MMRV3 struct {
	Account struct {
		Puuid string `json:"puuid"`
		Name  string `json:"name"`
		Tag   string `json:"tag"`
	} `json:"account"`
	Peak struct {
		Season struct {
			ID    string `json:"id"`
			Short string `json:"short"`
		} `json:"season"`
		RankingSchema string `json:"ranking_schema"`
		Tier          struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"tier"`
	} `json:"peak"`
	Current struct {
		Tier struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"tier"`
		Rr                   int `json:"rr"`
		LastChange           int `json:"last_change"`
		Elo                  int `json:"elo"`
		GamesNeededForRating int `json:"games_needed_for_rating"`
		LeaderboardPlacement struct {
			Rank      int       `json:"rank"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"leaderboard_placement"`
	} `json:"current"`
	Seasonal []struct {
		Season struct {
			ID    string `json:"id"`
			Short string `json:"short"`
		} `json:"season"`
		Wins    int `json:"wins"`
		Games   int `json:"games"`
		EndTier struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"end_tier"`
		RankingSchema        string `json:"ranking_schema"`
		LeaderboardPlacement struct {
			Rank      int       `json:"rank"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"leaderboard_placement"`
		ActWins []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"act_wins"`
	} `json:"seasonal"`
}
