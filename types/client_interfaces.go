package types

import (
	"context"
	"github.com/lanxre/go-valo/types/henrik_responses/accounts"
	"github.com/lanxre/go-valo/types/henrik_responses/content"
	"github.com/lanxre/go-valo/types/henrik_responses/esports"
	"github.com/lanxre/go-valo/types/henrik_responses/leaderboard"
	"github.com/lanxre/go-valo/types/henrik_responses/match"
	"github.com/lanxre/go-valo/types/henrik_responses/matchlist"
	"github.com/lanxre/go-valo/types/henrik_responses/mmr"
	"github.com/lanxre/go-valo/types/henrik_responses/mmr_history"
	"github.com/lanxre/go-valo/types/henrik_responses/stored_matches"
	"github.com/lanxre/go-valo/types/henrik_responses/website"
)

type Client interface {
	GetRaw(path string) ([]byte, error)
	GetRawWithContext(ctx context.Context, path string) ([]byte, error)
}

type HenrikClient interface {
	// account methods
	GetAccountV1(params ValorantAccountParams, options ...map[string]string) (accounts.AccountV1, error)
	GetAccountV2(params ValorantAccountParams, options ...map[string]string) (accounts.AccountV2, error)
	GetAccountByPUUIDV1(params PlayerParams, options ...map[string]string) (accounts.AccountV1, error)
	GetAccountByPUUIDV2(params PlayerParams, options ...map[string]string) (accounts.AccountV1, error)

	// content method
	GetContent(options ...map[string]string) (content.Content, error)

	// croshair method
	GetCroshair(options ...map[string]string) ([]byte, error)

	// esports schedule method
	GetScheduleEsports(options ...map[string]string) (esports.ScheduleEsports, error)

	// leaderboard method
	GetLeaderboard(params LeaderboardParams, options ...map[string]string) (leaderboard.Leaderboard, error)

	// matchlist methods
	GetMatchesV3(params MatchlistParamV3, options ...map[string]string) (matchlist.MatchesV3, error)
	GetMatchesV4(params MatchlistParamV4, options ...map[string]string) (matchlist.MatchesV4, error)
	GetMatchesByPUUIDV3(params PlayerRegionParams, options ...map[string]string) (matchlist.MatchesV3, error)
	GetMatchesByPUUIDV4(params PlayerRegionPlatformParams, options ...map[string]string) (matchlist.MatchesV4, error)

	// match methods
	GetMatchV2(params MatchParamV2, options ...map[string]string) (match.MatchV2, error)
	GetMatchV4(params MatchParamV4, options ...map[string]string) (match.MatchV4, error)

	// mmr history methods
	GetMMRHistoryV1(params MatchlistParamV3, options ...map[string]string) (mmrhistory.MMRHistoryV1, error)
	GetMMRHistoryByPUUIDV1(params PlayerRegionParams, options ...map[string]string) (mmrhistory.MMRHistoryV1, error)
	GetMMRHistoryV2(params MatchlistParamV4, options ...map[string]string) (mmrhistory.MMRHistoryV2, error)
	GetMMRHistoryByPUUIDV2(params PlayerRegionPlatformParams, options ...map[string]string) (mmrhistory.MMRHistoryV2, error)

	// mmr methods
	GetMMRV2(params MatchlistParamV3, options ...map[string]string) (mmr.MMRV2, error)
	GetMMRByPUUIDV2(params PlayerRegionParams, options ...map[string]string) (mmr.MMRV2, error)
	GetMMRV3(params MatchlistParamV4, options ...map[string]string) (mmr.MMRV3, error)
	GetMMRByPUUIDV3(params PlayerRegionPlatformParams, options ...map[string]string) (mmr.MMRV3, error)

	// stored data methods
	GetStoredMatchesV1(params MatchlistParamV3, options ...map[string]string) (storedmatches.StoredMatchesV1, error)
	GetStoredMatchesPUUIDV1(params PlayerRegionParams, options ...map[string]string) (storedmatches.StoredMatchesV1, error)
	GetStoredMMRHistoryV1(params MatchlistParamV3, options ...map[string]string) (storedmatches.StoredMMRHistoryV1, error)
	GetStoredMMRHistoryV2(params MatchlistParamV4, options ...map[string]string) (storedmatches.StoredMMRHistoryV2, error)
	GetStoredMMRHistoryByPUUIDV1(params PlayerRegionParams, options ...map[string]string) (storedmatches.StoredMMRHistoryV1, error)
	GetStoredMMRHistoryByPUUIDV2(params PlayerRegionPlatformParams, options ...map[string]string) (storedmatches.StoredMMRHistoryV2, error)

	// website method
	GetWebsiteNews(params WebsiteNewsParams, options ...map[string]string) (website.WebsiteNews, error)
}
