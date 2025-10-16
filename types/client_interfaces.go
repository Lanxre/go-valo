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
	"github.com/lanxre/go-valo/types/henrik_responses/website"
    "github.com/lanxre/go-valo/types/henrik_responses/stored_matches"
)

type Client interface {
	GetRaw(path string) ([]byte, error)
	GetRawWithContext(ctx context.Context, path string) ([]byte, error)
}

type HenrikClient interface {
	// account methods
	GetAccountV1(options ...map[string]string) (accounts.AccountV1, error)
	GetAccountV2(options ...map[string]string) (accounts.AccountV2, error)
	GetAccountByPUUIDV1(options ...map[string]string) (accounts.AccountV1, error)
	GetAccountByPUUIDV2(options ...map[string]string) (accounts.AccountV1, error)

	// content method
	GetContent(options ...map[string]string) (content.Content, error)

	// croshair method
	GetCroshair(options ...map[string]string) ([]byte, error)

	// esports schedule method
	GetScheduleEsports(options ...map[string]string) (esports.ScheduleEsports, error)

	// leaderboard method
	GetLeaderboard(options ...map[string]string) (leaderboard.Leaderboard, error)

	// matchlist methods
	GetMatchesV3(options ...map[string]string) (matchlist.MatchesV3, error)
	GetMatchesV4(options ...map[string]string) (matchlist.MatchesV4, error)
	GetMatchesByPUUIDV3(options ...map[string]string) (matchlist.MatchesV3, error)
	GetMatchesByPUUIDV4(options ...map[string]string) (matchlist.MatchesV4, error)

	// match methods
	GetMatchV2(options ...map[string]string) (match.MatchV2, error)
	GetMatchV4(options ...map[string]string) (match.MatchV4, error)

	// mmr history methods
	GetMMRHistoryV1(options ...map[string]string) (mmrhistory.MMRHistoryV1, error)
	GetMMRHistoryByPUUIDV1(options ...map[string]string) (mmrhistory.MMRHistoryV1, error)
	GetMMRHistoryV2(options ...map[string]string) (mmrhistory.MMRHistoryV2, error)
	GetMMRHistoryByPUUIDV2(options ...map[string]string) (mmrhistory.MMRHistoryV2, error)

	// mmr methods
	GetMMRV2(options ...map[string]string) (mmr.MMRV2, error)
	GetMMRByPUUIDV2(options ...map[string]string) (mmr.MMRV2, error)
	GetMMRV3(options ...map[string]string) (mmr.MMRV3, error)
	GetMMRByPUUIDV3(options ...map[string]string) (mmr.MMRV3, error)

	// stored data methods
	GetStoredMatchesV1(options ...map[string]string) (storedmatches.StoredMatchesV1, error)
	GetStoredMatchesPUUIDV1(options ...map[string]string) (storedmatches.StoredMatchesV1, error)
	GetStoredMMRHistoryV1(options ...map[string]string) (storedmatches.StoredMMRHistoryV1, error)
	GetStoredMMRHistoryV2(options ...map[string]string) (storedmatches.StoredMMRHistoryV2, error)
	GetStoredMMRHistoryByPUUIDV1(options ...map[string]string) (storedmatches.StoredMMRHistoryV1, error)
	GetStoredMMRHistoryByPUUIDV2(options ...map[string]string) (storedmatches.StoredMMRHistoryV2, error)

	// website method
	GetWebsiteNews(options ...map[string]string) (website.WebsiteNews, error)
}
