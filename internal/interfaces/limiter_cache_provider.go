package interfaces

import "context"

type LimiterCacheProvider interface {
	GetLimitInfo(ctx context.Context, key string) (int, error)
	SetLimitInfo(ctx context.Context, key string) error
	BlockCheck(ctx context.Context, key string) (bool, error)
	Block(ctx context.Context, key string, blockTimeInSeconds int) error
}
