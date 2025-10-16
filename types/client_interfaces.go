package types

import (
	"context"
	"github.com/lanxre/go-valo/types/henrik_responses/accounts"
    "github.com/lanxre/go-valo/types/henrik_responses/content"
    "github.com/lanxre/go-valo/types/henrik_responses/esports"
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
}