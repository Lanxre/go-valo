package storedmatches

import "time"

type StoredMMRHistoryV1 []struct {
	MatchID string `json:"match_id"`
	Tier    struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"tier"`
	Map struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"map"`
	Season struct {
		ID    string `json:"id"`
		Short string `json:"short"`
	} `json:"season"`
	RankingInTier int       `json:"ranking_in_tier"`
	LastMmrChange int       `json:"last_mmr_change"`
	Elo           int       `json:"elo"`
	Date          time.Time `json:"date"`
}