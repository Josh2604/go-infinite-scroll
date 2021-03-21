package providers

import (
	"context"

	"github.com/Josh2604/go-infinite-scroll/api/core/contracts/getposts"
)

type Posts interface {
	GetPosts(ctx context.Context, paginator *getposts.Paginator) (*getposts.Response, error)
}
