package mmr

type MMRV2 struct {
	Name        string `json:"name"`
	Tag         string `json:"tag"`
	CurrentData struct {
		Currenttier        int    `json:"currenttier"`
		CurrenttierPatched string `json:"currenttier_patched"`
		Images             struct {
			Small        string `json:"small"`
			Large        string `json:"large"`
			TriangleDown string `json:"triangle_down"`
			TriangleUp   string `json:"triangle_up"`
		} `json:"images"`
		RankingInTier       int  `json:"ranking_in_tier"`
		MmrChangeToLastGame int  `json:"mmr_change_to_last_game"`
		Elo                 int  `json:"elo"`
		Old                 bool `json:"old"`
	} `json:"current_data"`
	HighestRank struct {
		Old         bool   `json:"old"`
		Tier        int    `json:"tier"`
		PatchedTier string `json:"patched_tier"`
		Season      string `json:"season"`
	} `json:"highest_rank"`
	BySeason struct {
		E6A3 struct {
			Error            bool   `json:"error"`
			Wins             int    `json:"wins"`
			NumberOfGames    int    `json:"number_of_games"`
			FinalRank        int    `json:"final_rank"`
			FinalRankPatched string `json:"final_rank_patched"`
			ActRankWins      []struct {
				PatchedTier string `json:"patched_tier"`
				Tier        int    `json:"tier"`
			} `json:"act_rank_wins"`
			Old bool `json:"old"`
		} `json:"e6a3"`
		E6A2 struct {
			Error            bool   `json:"error"`
			Wins             int    `json:"wins"`
			NumberOfGames    int    `json:"number_of_games"`
			FinalRank        int    `json:"final_rank"`
			FinalRankPatched string `json:"final_rank_patched"`
			ActRankWins      []struct {
				PatchedTier string `json:"patched_tier"`
				Tier        int    `json:"tier"`
			} `json:"act_rank_wins"`
			Old bool `json:"old"`
		} `json:"e6a2"`
		E6A1 struct {
			Error            bool   `json:"error"`
			Wins             int    `json:"wins"`
			NumberOfGames    int    `json:"number_of_games"`
			FinalRank        int    `json:"final_rank"`
			FinalRankPatched string `json:"final_rank_patched"`
			ActRankWins      []struct {
				PatchedTier string `json:"patched_tier"`
				Tier        int    `json:"tier"`
			} `json:"act_rank_wins"`
			Old bool `json:"old"`
		} `json:"e6a1"`
		E5A3 struct {
			Error            bool   `json:"error"`
			Wins             int    `json:"wins"`
			NumberOfGames    int    `json:"number_of_games"`
			FinalRank        int    `json:"final_rank"`
			FinalRankPatched string `json:"final_rank_patched"`
			ActRankWins      []struct {
				PatchedTier string `json:"patched_tier"`
				Tier        int    `json:"tier"`
			} `json:"act_rank_wins"`
			Old bool `json:"old"`
		} `json:"e5a3"`
		E5A2 struct {
			Error            bool   `json:"error"`
			Wins             int    `json:"wins"`
			NumberOfGames    int    `json:"number_of_games"`
			FinalRank        int    `json:"final_rank"`
			FinalRankPatched string `json:"final_rank_patched"`
			ActRankWins      []struct {
				PatchedTier string `json:"patched_tier"`
				Tier        int    `json:"tier"`
			} `json:"act_rank_wins"`
			Old bool `json:"old"`
		} `json:"e5a2"`
		E5A1 struct {
			Error            bool   `json:"error"`
			Wins             int    `json:"wins"`
			NumberOfGames    int    `json:"number_of_games"`
			FinalRank        int    `json:"final_rank"`
			FinalRankPatched string `json:"final_rank_patched"`
			ActRankWins      []struct {
				PatchedTier string `json:"patched_tier"`
				Tier        int    `json:"tier"`
			} `json:"act_rank_wins"`
			Old bool `json:"old"`
		} `json:"e5a1"`
		E4A3 struct {
			Error            bool   `json:"error"`
			Wins             int    `json:"wins"`
			NumberOfGames    int    `json:"number_of_games"`
			FinalRank        int    `json:"final_rank"`
			FinalRankPatched string `json:"final_rank_patched"`
			ActRankWins      []struct {
				PatchedTier string `json:"patched_tier"`
				Tier        int    `json:"tier"`
			} `json:"act_rank_wins"`
			Old bool `json:"old"`
		} `json:"e4a3"`
		E4A2 struct {
			Error            bool   `json:"error"`
			Wins             int    `json:"wins"`
			NumberOfGames    int    `json:"number_of_games"`
			FinalRank        int    `json:"final_rank"`
			FinalRankPatched string `json:"final_rank_patched"`
			ActRankWins      []struct {
				PatchedTier string `json:"patched_tier"`
				Tier        int    `json:"tier"`
			} `json:"act_rank_wins"`
			Old bool `json:"old"`
		} `json:"e4a2"`
		E4A1 struct {
			Error            bool   `json:"error"`
			Wins             int    `json:"wins"`
			NumberOfGames    int    `json:"number_of_games"`
			FinalRank        int    `json:"final_rank"`
			FinalRankPatched string `json:"final_rank_patched"`
			ActRankWins      []struct {
				PatchedTier string `json:"patched_tier"`
				Tier        int    `json:"tier"`
			} `json:"act_rank_wins"`
			Old bool `json:"old"`
		} `json:"e4a1"`
		E3A3 struct {
			Error            bool   `json:"error"`
			Wins             int    `json:"wins"`
			NumberOfGames    int    `json:"number_of_games"`
			FinalRank        int    `json:"final_rank"`
			FinalRankPatched string `json:"final_rank_patched"`
			ActRankWins      []struct {
				PatchedTier string `json:"patched_tier"`
				Tier        int    `json:"tier"`
			} `json:"act_rank_wins"`
			Old bool `json:"old"`
		} `json:"e3a3"`
		E3A2 struct {
			Error            bool   `json:"error"`
			Wins             int    `json:"wins"`
			NumberOfGames    int    `json:"number_of_games"`
			FinalRank        int    `json:"final_rank"`
			FinalRankPatched string `json:"final_rank_patched"`
			ActRankWins      []struct {
				PatchedTier string `json:"patched_tier"`
				Tier        int    `json:"tier"`
			} `json:"act_rank_wins"`
			Old bool `json:"old"`
		} `json:"e3a2"`
		E3A1 struct {
			Error            bool   `json:"error"`
			Wins             int    `json:"wins"`
			NumberOfGames    int    `json:"number_of_games"`
			FinalRank        int    `json:"final_rank"`
			FinalRankPatched string `json:"final_rank_patched"`
			ActRankWins      []struct {
				PatchedTier string `json:"patched_tier"`
				Tier        int    `json:"tier"`
			} `json:"act_rank_wins"`
			Old bool `json:"old"`
		} `json:"e3a1"`
		E2A3 struct {
			Error            bool   `json:"error"`
			Wins             int    `json:"wins"`
			NumberOfGames    int    `json:"number_of_games"`
			FinalRank        int    `json:"final_rank"`
			FinalRankPatched string `json:"final_rank_patched"`
			ActRankWins      []struct {
				PatchedTier string `json:"patched_tier"`
				Tier        int    `json:"tier"`
			} `json:"act_rank_wins"`
			Old bool `json:"old"`
		} `json:"e2a3"`
		E2A2 struct {
			Error            bool   `json:"error"`
			Wins             int    `json:"wins"`
			NumberOfGames    int    `json:"number_of_games"`
			FinalRank        int    `json:"final_rank"`
			FinalRankPatched string `json:"final_rank_patched"`
			ActRankWins      []struct {
				PatchedTier string `json:"patched_tier"`
				Tier        int    `json:"tier"`
			} `json:"act_rank_wins"`
			Old bool `json:"old"`
		} `json:"e2a2"`
		E2A1 struct {
			Error            bool   `json:"error"`
			Wins             int    `json:"wins"`
			NumberOfGames    int    `json:"number_of_games"`
			FinalRank        int    `json:"final_rank"`
			FinalRankPatched string `json:"final_rank_patched"`
			ActRankWins      []struct {
				PatchedTier string `json:"patched_tier"`
				Tier        int    `json:"tier"`
			} `json:"act_rank_wins"`
			Old bool `json:"old"`
		} `json:"e2a1"`
		E1A3 struct {
			Error            bool   `json:"error"`
			Wins             int    `json:"wins"`
			NumberOfGames    int    `json:"number_of_games"`
			FinalRank        int    `json:"final_rank"`
			FinalRankPatched string `json:"final_rank_patched"`
			ActRankWins      []struct {
				PatchedTier string `json:"patched_tier"`
				Tier        int    `json:"tier"`
			} `json:"act_rank_wins"`
			Old bool `json:"old"`
		} `json:"e1a3"`
		E1A2 struct {
			Error            bool   `json:"error"`
			Wins             int    `json:"wins"`
			NumberOfGames    int    `json:"number_of_games"`
			FinalRank        int    `json:"final_rank"`
			FinalRankPatched string `json:"final_rank_patched"`
			ActRankWins      []struct {
				PatchedTier string `json:"patched_tier"`
				Tier        int    `json:"tier"`
			} `json:"act_rank_wins"`
			Old bool `json:"old"`
		} `json:"e1a2"`
		E1A1 struct {
			Error            bool   `json:"error"`
			Wins             int    `json:"wins"`
			NumberOfGames    int    `json:"number_of_games"`
			FinalRank        int    `json:"final_rank"`
			FinalRankPatched string `json:"final_rank_patched"`
			ActRankWins      []struct {
				PatchedTier string `json:"patched_tier"`
				Tier        int    `json:"tier"`
			} `json:"act_rank_wins"`
			Old bool `json:"old"`
		} `json:"e1a1"`
	} `json:"by_season"`
}
