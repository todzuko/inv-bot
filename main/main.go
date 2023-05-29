package main

import (
	"github.com/joho/godotenv"
	"github.com/todzuko/inv-bot/telegram"
)

func main() {
	loadEnv()
	telegram.Connect()
}

func loadEnv() {
	err := godotenv.Load("/data/.env")
	if err != nil {
		panic("err loading: " + err.Error())
	}
}
