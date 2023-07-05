package handlers

import (
	"fmt"
	"github.com/todzuko/inv-bot/api/database"
	"os"
	"strconv"
)

func SetLimit(chat int64, inputLimit string) string {
	if isValidLimit(inputLimit) {
		numLimit, _ := strconv.Atoi(inputLimit)
		ok := database.SetLimit(chat, numLimit)
		if !ok {
			return "Возникла проблема с установкой липита, попробуй еще раз /setLimit"
		}
		return fmt.Sprintf("Лимит %d установлен\nВы можете проверить лимит, используя команду /getLimit", numLimit)
	}
	return fmt.Sprintf("Неверный лимит %s, введите целое число больше или равное нулю", inputLimit)
}

func GetLimit(chat int64) string {
	limit, ok := database.GetLimit(chat)
	if ok {
		return fmt.Sprintf("Текущий лимит: %d", limit)
	} else {
		return fmt.Sprintf("Лимит не установлен\nЛимит по умолчанию - %s\nИспользуйте команду /setLimit, чтобы задать лимит", os.Getenv("DEFAULT_LIMIT"))
	}
}

func isValidLimit(num string) bool {
	return true
}
