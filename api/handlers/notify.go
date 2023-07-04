package handlers

import (
	"github.com/todzuko/inv-bot/api/external/tinkoff-api"
	"github.com/todzuko/inv-bot/api/helpers"
)

func Notify() string {
	stocks := tinkoff_api.Instr()
	return helpers.ConstructReport(stocks)
}
