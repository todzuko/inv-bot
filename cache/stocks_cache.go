package cache

import (
	"encoding/json"
	"fmt"
	tinkoff_api "github.com/todzuko/inv-bot/api/external/tinkoff-api"
	"strconv"
	"strings"
	"time"
)

type DividendStocksCache interface {
	Set(key string, value *[]tinkoff_api.DividendStock)
	Get(key string) *[]tinkoff_api.DividendStock
	GetStocksWithLimit(limit int) *[]tinkoff_api.DividendStock
}

func (cache *redisCache) Set(key string, value *[]tinkoff_api.DividendStock) {
	client := cache.getClient()
	jsonVal, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	client.Set(ctx, key, jsonVal, cache.expires*time.Second)
}

func (cache *redisCache) Get(key string) *[]tinkoff_api.DividendStock {
	client := cache.getClient()
	jsonVal, err := client.Get(ctx, key).Result()
	if err != nil {
		limit := getLimitFromKey(key)
		if limit < 0 {
			return nil
		}
		stocks := cache.SetStocksWithLimit(limit)
		return stocks
	}
	stocks := []tinkoff_api.DividendStock{}
	err = json.Unmarshal([]byte(jsonVal), &stocks)
	if err != nil {
		panic(err)
	}
	return &stocks
}

func (cache *redisCache) SetStocksWithLimit(limit int) *[]tinkoff_api.DividendStock {
	stocks := tinkoff_api.GetDividendStocks(limit)
	key := getLimitKey(limit)
	cache.Set(key, stocks)
	return stocks
}

func (cache *redisCache) GetStocksWithLimit(limit int) *[]tinkoff_api.DividendStock {
	key := getLimitKey(limit)
	return cache.Get(key)
}

func getLimitKey(limit int) string {
	return fmt.Sprintf("stocks_with_limit_%d", limit)
}

func getLimitFromKey(key string) int {
	keyParts := strings.Split(key, "_")
	limitStr := keyParts[len(keyParts)-1]
	limit, err := strconv.Atoi(limitStr)
	if err == nil {
		return limit
	}
	return -1
}
