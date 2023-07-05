package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/todzuko/inv-bot/api/database"
	"github.com/todzuko/inv-bot/telegram"
)

func main() {
	loadEnv()
	dbClient, err := database.Connect()
	if err != nil {
		panic("err loading: " + err.Error())
	}
	telegram.Connect()
	defer dbClient.Disconnect(context.Background())
}

func loadEnv() {
	err := godotenv.Load("/data/.env")
	if err != nil {
		panic("err loading: " + err.Error())
	}
}
