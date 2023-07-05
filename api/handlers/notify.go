package handlers

import (
	"github.com/todzuko/inv-bot/api/database"
	"github.com/todzuko/inv-bot/api/helpers"
	"os"
	"strconv"
)

func Notify(chat int64) string {
	limit := getLimitByChat(chat)
	stocks := helpers.GetDividendStocksWithProfitLimit(limit)
	return helpers.ConstructReport(stocks)
}

func getLimitByChat(chatId int64) int {
	limit, ok := database.GetLimit(chatId)
	if !ok {
		limit, _ = strconv.Atoi(os.Getenv("DEFAULT_LIMIT"))
	}
	return limit
}
