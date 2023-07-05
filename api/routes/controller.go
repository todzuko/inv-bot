package routes

import (
	"fmt"
	"github.com/todzuko/inv-bot/api/handlers"
	"strconv"
	"strings"
)

func GetResponse(path string, chatId int64) string {
	path = stripPathName(path)
	switch path {
	case `help`:
		fallthrough
	case `start`:
		return handlers.Help()
	case `getLimit`:
		return handlers.GetLimit(chatId)
	case `setLimit`:
		return "Введи целое число >= 0"
	case `notify`:
		return handlers.Notify(chatId)
	default:
		if _, err := strconv.Atoi(path); err == nil {
			return handlers.SetLimit(chatId, path)
		}
		return "Команды \"" + path + "\" не существует"
	}
}

func stripPathName(path string) string {
	if string(path[0]) == `/` {
		path = strings.Replace(path, "/", "", 1)
	}
	fmt.Println(path)
	return path
}
