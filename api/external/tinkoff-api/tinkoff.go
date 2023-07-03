package tinkoff_api

import (
	"context"
	"github.com/tinkoff/invest-api-go-sdk/investgo"
	"go.uber.org/zap"
	"log"
)

func GetClient() (*investgo.Client, *zap.SugaredLogger, context.CancelFunc) {
	config, err := investgo.LoadConfig("/data/config.yaml")
	if err != nil {
		log.Fatalf("config loading error %v", err.Error())
	}
	ctx, cancel := context.WithCancel(context.Background())

	prod, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("logger creating error %e", err)
	}
	logger := prod.Sugar()

	client, err := investgo.NewClient(ctx, config, logger)
	if err != nil {
		logger.Fatalf("Client creating error %v", err.Error())
	}

	err = prod.Sync()
	return client, logger, cancel
}
