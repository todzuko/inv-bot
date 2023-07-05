package tinkoff_api

import (
	"fmt"
	"github.com/tinkoff/invest-api-go-sdk/investgo"
	investapi "github.com/tinkoff/invest-api-go-sdk/proto"
	"github.com/todzuko/inv-bot/api/database"
	"go.uber.org/zap"
	"os"
	"strconv"
	"time"
)

type DividendStock struct {
	Name           string
	Code           string
	DividendDate   time.Time
	DividendAmount *investapi.MoneyValue
	DividendProfit *investapi.Quotation
}

func Instr(chat int64) *[]DividendStock {
	client, logger, cancel := GetClient()
	var divStocks []DividendStock
	instrumentsService := client.NewInstrumentsServiceClient()
	instrResp, err := instrumentsService.Shares(1)
	if err != nil {
		logger.Errorf(err.Error())
		return &divStocks
	}
	ins := instrResp.GetInstruments()
	var divs *investgo.GetDividendsResponse
	for _, instrument := range ins {
		if instrument.Currency == "rub" && instrument.DivYieldFlag == true && instrument.LiquidityFlag == true {
			divs, err = instrumentsService.GetDividents(instrument.Figi, time.Now(), time.Now().AddDate(0, 6, 0))
			if err != nil {
				fmt.Println(instrument.Name, err)
			}
			if err == nil && len(divs.Dividends) > 0 {
				processDividendStock(instrument, divs, &divStocks, chat)
			}
		}
	}
	defer stopClient(client, logger)
	defer cancel()
	return &divStocks
}

func processDividendStock(instrument *investapi.Share, divs *investgo.GetDividendsResponse, divStocks *[]DividendStock, chat int64) {
	closestDiv := divs.Dividends[0]
	limit, ok := database.GetLimit(chat)
	if !ok {
		limit, _ = strconv.Atoi(os.Getenv("DEFAULT_LIMIT"))
	}
	if closestDiv.DividendNet.Currency == "rub" && closestDiv.YieldValue != nil && closestDiv.YieldValue.Units > int64(limit) {
		stockName := instrument.GetName()
		stockCode := instrument.Ticker
		divDate := closestDiv.PaymentDate.AsTime()
		divAmount := closestDiv.DividendNet
		divProfit := closestDiv.YieldValue
		*divStocks = append(*divStocks, DividendStock{
			stockName,
			stockCode,
			divDate,
			divAmount,
			divProfit,
		})
	}
}

func stopClient(client *investgo.Client, logger *zap.SugaredLogger) {
	logger.Infof("Closing client connection")
	err := client.Stop()
	if err != nil {
		logger.Error("client shutdown error %v", err.Error())
	}
}
