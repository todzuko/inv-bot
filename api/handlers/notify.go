package handlers

import (
	"github.com/todzuko/inv-bot/api/external/tinkoff-api"
	"github.com/todzuko/inv-bot/api/helpers"
)

func Notify(chat int64) string {
	stocks := tinkoff_api.Instr(chat)
	return helpers.ConstructReport(stocks)
}
