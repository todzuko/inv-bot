package routes

import (
	"fmt"
	"github.com/todzuko/inv-bot/api/handlers"
	"strings"
)

func GetResponse(path string) string {
	path = stripPathName(path)
	switch path {
	case `start`:
		return handlers.Start()
	case `getLimit`:
		return handlers.GetLimit()
	case `setLimit`:
		return handlers.SetLimit()
	case `notify`:
		return handlers.Notify()
	default:
		return "Command \"" + path + "\" does not exist"
	}
}

func stripPathName(path string) string {
	if string(path[0]) == `/` {
		path = strings.Replace(path, "/", "", 1)
	}
	fmt.Println(path)
	return path
}
