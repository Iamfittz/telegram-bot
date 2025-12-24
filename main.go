package main

import (
	"context"
	"log"

	"github.com/iamfittz/telegram-bot/bot"
	"github.com/iamfittz/telegram-bot/telemetry"
)

func main() {
	ctx := context.Background()

	cleanup, err := telemetry.InitTelemetry(ctx, "telegram-bot")
	if err != nil {
		log.Fatalf("Failed to init telemetry: %v", err)
	}
	defer cleanup()

	bot.Start(ctx)
}
