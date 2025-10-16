package mmrhistory

type MMRHistoryV1 []struct {
	Currenttier        int    `json:"currenttier"`
	CurrenttierPatched string `json:"currenttier_patched"`
	Images             struct {
		Small        string `json:"small"`
		Large        string `json:"large"`
		TriangleDown string `json:"triangle_down"`
		TriangleUp   string `json:"triangle_up"`
	} `json:"images"`
	MatchID string `json:"match_id"`
	Map     struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"map"`
	SeasonID            string `json:"season_id"`
	RankingInTier       int    `json:"ranking_in_tier"`
	MmrChangeToLastGame int    `json:"mmr_change_to_last_game"`
	Elo                 int    `json:"elo"`
	Date                string `json:"date"`
	DateRaw             int    `json:"date_raw"`
}
