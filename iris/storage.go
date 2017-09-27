package iris

import (
	"github.com/kataras/iris/context"
	"github.com/xuebing1110/fortify/user/storage"
	"github.com/xuebing1110/fortify/user/storage/redis"
)

const (
	CONTEXT_STORAGE_TAG = "Storage"
)

func UseRedisStorage(ctx context.Context) {
	var redis_store storage.Storage = redis.Client
	ctx.Values().Set(CONTEXT_STORAGE_TAG, redis_store)

	ctx.Next()
}

func GetStorage(ctx context.Context) (store storage.Storage, ok bool) {
	store, ok = ctx.Values().Get(CONTEXT_STORAGE_TAG).(storage.Storage)
	return store, ok
}
