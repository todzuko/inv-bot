package helpers

import (
	"bytes"
	"fmt"
	tinkoff_api "github.com/todzuko/inv-bot/api/external/tinkoff-api"
	"strconv"
)

func ConstructReport(divStocks *[]tinkoff_api.DividendStock) string {
	if len(*divStocks) < 1 {
		return "Подходящие акции не найдены"
	}
	var buffer bytes.Buffer
	var day int
	var month string
	var amount string
	var profit string

	for _, stock := range *divStocks {
		buffer.WriteString(fmt.Sprintf("▫️%s (%s)\n", stock.Name, stock.Code))

		day = stock.DividendDate.Day()
		month = stock.DividendDate.Month().String()
		month = GetMonthRu(month)
		if month == "" {
			fmt.Println("Invalid month")
			return ""
		}
		buffer.WriteString(strconv.Itoa(day))
		buffer.WriteString(" ")
		buffer.WriteString(month)
		buffer.WriteString(" | ")

		amount = formatAmount(stock.DividendAmount.Units, stock.DividendAmount.Nano)
		buffer.WriteString(amount)
		buffer.WriteString("₽ | ")

		profit = formatAmount(stock.DividendProfit.Units, stock.DividendProfit.Nano)
		buffer.WriteString(profit)
		buffer.WriteString("%\n")
	}
	return buffer.String()
}

func formatAmount(amount int64, amountReminder int32) string {
	amountStr := strconv.FormatInt(amount, 10)
	amountReminderStr := strconv.Itoa(int(amountReminder))
	if len(amountReminderStr) > 2 {
		amountReminderStr = amountReminderStr[:2]
	}
	return fmt.Sprintf("%s.%s", amountStr, amountReminderStr)
}
