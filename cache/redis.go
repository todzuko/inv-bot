package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

var ctx context.Context

func NewRedisCache(host string, db int, exp time.Duration) DividendStocksCache {
	ctx = context.Background()
	return &redisCache{
		host,
		db,
		exp,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}
