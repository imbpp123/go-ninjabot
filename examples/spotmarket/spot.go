package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/imbpp123/go-ninjabot"
	"github.com/imbpp123/go-ninjabot/examples/strategies"
	"github.com/imbpp123/go-ninjabot/exchange"
)

// This example shows how to use spot market with NinjaBot in Binance
func main() {
	var (
		ctx             = context.Background()
		apiKey          = os.Getenv("API_KEY")
		secretKey       = os.Getenv("API_SECRET")
		telegramToken   = os.Getenv("TELEGRAM_TOKEN")
		telegramUser, _ = strconv.Atoi(os.Getenv("TELEGRAM_USER"))
	)

	settings := ninjabot.Settings{
		Pairs: []string{
			"BTCUSDT",
			"ETHUSDT",
		},
		Telegram: ninjabot.TelegramSettings{
			Enabled: true,
			Token:   telegramToken,
			Users:   []int{telegramUser},
		},
	}

	// Initialize your exchange
	binance, err := exchange.NewBinance(ctx, exchange.WithBinanceCredentials(apiKey, secretKey))
	if err != nil {
		log.Fatalln(err)
	}

	// Initialize your strategy and bot
	strategy := new(strategies.CrossEMA)
	bot, err := ninjabot.NewBot(ctx, settings, binance, strategy)
	if err != nil {
		log.Fatalln(err)
	}

	err = bot.Run(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}
