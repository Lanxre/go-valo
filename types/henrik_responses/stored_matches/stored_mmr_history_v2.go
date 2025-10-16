package storedmatches

import "time"

type StoredMMRHistoryV2 []struct {
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
	Rr         int       `json:"rr"`
	LastChange int       `json:"last_change"`
	Elo        int       `json:"elo"`
	Date       time.Time `json:"date"`
}