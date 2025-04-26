package callback

import (
	"context"
	"telegarm/config"
	"telegarm/handlers"
	"telegarm/handlers/input"
	"telegarm/storage"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func HandlerBuy(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})
	PayUrl := handlers.CreatePayURL("connect", 157570)

	PatternName := update.CallbackQuery.Data

	if PayUrl.Response.URL != "" {
		b.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
			MessageID: update.CallbackQuery.Message.Message.ID,
			Text:      "*" + bot.EscapeMarkdown("Покупка") + "*" + bot.EscapeMarkdown(" Напишите ник на который покупаете данный товар, ожидание 30с"),
			ParseMode: models.ParseModeMarkdown,
		})

		inputChan := make(chan string)
		storage.UserInfoChat.Store(update.CallbackQuery.From.ID, inputChan)
		go input.WaitNickNameInput(ctx, b, update.CallbackQuery.Message.Message.Chat.ID, update.CallbackQuery.From.ID, inputChan, PatternName, config.ProductID[PatternName])
	} else {
		kb := &models.InlineKeyboardMarkup{
			InlineKeyboard: [][]models.InlineKeyboardButton{
				{
					{Text: "Повторить попытку: " + config.PriceProduct[PatternName] + "₽", CallbackData: PatternName},
					{Text: "Меню", CallbackData: "menu"},
				}, {},
			},
		}

		b.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
			MessageID:   update.CallbackQuery.Message.Message.ID,
			Text:        "*" + bot.EscapeMarkdown("Ошибка") + "*" + bot.EscapeMarkdown("\nК сожалению, в данный момент оплата недоступна из-за технических неполадок. Мы уже работаем над устранением проблемы и скоро всё наладится. Пожалуйста, попробуйте позднее."),
			ParseMode:   models.ParseModeMarkdown,
			ReplyMarkup: kb,
		})
	}
}
