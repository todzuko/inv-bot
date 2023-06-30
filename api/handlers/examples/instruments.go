package examples

import (
	"bytes"
	"github.com/tinkoff/invest-api-go-sdk/investgo"
	tinkoff_api "github.com/todzuko/inv-bot/api/external/tinkoff-api"
	"strconv"
	"time"
)

func Instr() string {
	client, logger, cancel := tinkoff_api.GetClient()

	instrumentsService := client.NewInstrumentsServiceClient()
	instrResp, err := instrumentsService.Shares(1)
	if err != nil {
		logger.Errorf(err.Error())
	} else {
		var buffer bytes.Buffer
		ins := instrResp.GetInstruments()
		var divs *investgo.GetDividendsResponse
		for _, instrument := range ins {
			if instrument.Currency == "rub" {
				divs, err = instrumentsService.GetDividents(instrument.Figi, time.Now(), time.Now().AddDate(0, 6, 0))
				if err == nil && len(divs.Dividends) > 0 {
					closestDiv := divs.Dividends[0]
					if closestDiv.DividendNet.Currency == "rub" {
						if closestDiv.YieldValue != nil {
							if closestDiv.YieldValue.Units > 5 {
								buffer.WriteString("▫️")
								buffer.WriteString(instrument.GetName())
								buffer.WriteString(" (")
								buffer.WriteString(instrument.Ticker)
								buffer.WriteString(")")
								buffer.WriteString("\n")
								day := closestDiv.PaymentDate.AsTime().Day()
								month := closestDiv.PaymentDate.AsTime().Month().String()
								month = getRuMonth(month)
								buffer.WriteString(strconv.Itoa(day))
								buffer.WriteString(" ")
								buffer.WriteString(month)
								buffer.WriteString(" | ")
								amount := strconv.FormatInt(closestDiv.DividendNet.Units, 10)
								amountRem := strconv.Itoa(int(closestDiv.DividendNet.Nano))
								buffer.WriteString(amount)
								buffer.WriteString(".")
								if len(amountRem) > 2 {
									buffer.WriteString(amountRem[:2])
								} else {
									buffer.WriteString(amountRem)
								}
								buffer.WriteString("₽ | ")
								amount = strconv.FormatInt(closestDiv.YieldValue.Units, 10)
								amountRem = strconv.Itoa(int(closestDiv.YieldValue.Nano))
								buffer.WriteString(amount)
								buffer.WriteString(".")
								if len(amountRem) > 3 {
									buffer.WriteString(amountRem[:3])
								} else {
									buffer.WriteString(amountRem)
								}
								buffer.WriteString("%")

								buffer.WriteString("\n")
							}
						}
					}
				}
			}
		}
		return buffer.String()
	}
	defer func() {
		logger.Infof("Closing client connection")
		err := client.Stop()
		if err != nil {
			logger.Error("client shutdown error %v", err.Error())
		}
	}()
	defer cancel()
	return "not found"
}

func getRuMonth(month string) string {
	switch month {
	case "January":
		return "января"
	case "February":
		return "февраля"
	case "March":
		return "марта"
	case "April":
		return "апреля"
	case "May":
		return "мая"
	case "June":
		return "июня"
	case "July":
		return "июля"
	case "August":
		return "августа"
	case "September":
		return "сентября"
	case "October":
		return "октября"
	case "November":
		return "ноября"
	case "December":
		return "декабря"
	default:
		return ""
	}
}
