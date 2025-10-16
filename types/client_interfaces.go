package types

import (
	"context"
	"github.com/lanxre/go-valo/types/henrik_responses/accounts"
)

type Client interface {
    GetRaw(path string) ([]byte, error)
    GetRawWithContext(ctx context.Context, path string) ([]byte, error)
}

type HenrikClient interface {
    GetAccountV1(options ...map[string]string) (accounts.AccountV1, error)
}