package fulltextsearch

import (
	"context"

	"github.com/luongduc1246/ultility/reqparams"
)

type FullTextSearcher interface {
	CreateIndex(ctx context.Context, name string) error
	DeleteIndex(ctx context.Context, index string) error
	Insert(ctx context.Context, index string, id string, model interface{}) error
	Update(ctx context.Context, index string, id string, model interface{}) error
	Delete(ctx context.Context, index string, id string) error
	Get(ctx context.Context, index string, id string, model interface{}) error
	Search(ctx context.Context, param reqparams.Search, models interface{}) error
}
