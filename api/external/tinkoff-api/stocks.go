package tinkoff_api

import (
	"fmt"
	"github.com/tinkoff/invest-api-go-sdk/investgo"
	investapi "github.com/tinkoff/invest-api-go-sdk/proto"
	"go.uber.org/zap"
	"time"
)

type DividendStock struct {
	Name           string                `bson:"name"`
	Code           string                `bson:"code"`
	DividendDate   time.Time             `bson:"dividend_date"`
	DividendAmount *investapi.MoneyValue `bson:"dividend_amount"`
	DividendProfit *investapi.Quotation  `bson:"dividend_profit"`
}

func GetDividendStocks(limit int) *[]DividendStock {
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
				processDividendStock(instrument, divs, &divStocks, limit)
			}
		}
	}
	defer stopClient(client, logger)
	defer cancel()
	return &divStocks
}

func processDividendStock(instrument *investapi.Share, divs *investgo.GetDividendsResponse, divStocks *[]DividendStock, limit int) {
	closestDiv := divs.Dividends[0]
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
