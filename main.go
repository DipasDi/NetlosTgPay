package main

import (
	"context"
	"os"
	"os/signal"
	"telegarm/config"
	"telegarm/handlers/callback"
	"telegarm/handlers/input"

	"github.com/go-telegram/bot"
)

// Send any text message to the bot after the bot has been started

// Цены товаров

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(callback.StartHandeler),
		bot.WithCallbackQueryDataHandler("info", bot.MatchTypePrefix, callback.DefaultHandler),
		bot.WithCallbackQueryDataHandler("menu", bot.MatchTypePrefix, callback.DefaultHandler),
		bot.WithCallbackQueryDataHandler("other", bot.MatchTypePrefix, callback.DefaultHandler),
		bot.WithCallbackQueryDataHandler("buy", bot.MatchTypePrefix, callback.HandlerBuy),
	}

	b, err := bot.New(config.TokenBot, opts...)
	if nil != err {
		// panics for the sake of simplicity.
		// you should handle this error properly in your code.
		panic(err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypePrefix, bot.HandlerFunc(callback.StartHandeler))

	b.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypePrefix, bot.HandlerFunc(input.HandleNameInput))
	b.Start(ctx)
}
