package main

import (
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
}

func loadEnv() {
	err := godotenv.Load("/data/.env")
	if err != nil {
		panic("err loading: " + err.Error())
	}
}
