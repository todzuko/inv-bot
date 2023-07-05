package helpers

import (
	tinkoff_api "github.com/todzuko/inv-bot/api/external/tinkoff-api"
	"github.com/todzuko/inv-bot/cache"
	"os"
)

var stockCache = cache.NewRedisCache(
	os.Getenv("REDIS_HOST"),
	0,
	3600,
)

func GetDividendStocksWithProfitLimit(limit int) *[]tinkoff_api.DividendStock {
	stockList := stockCache.GetStocksWithLimit(limit)
	return stockList
}
