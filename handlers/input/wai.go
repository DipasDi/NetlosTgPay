package input

import (
	"context"
	"log"
	"telegarm/config"
	"telegarm/handlers"
	"telegarm/storage"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func WaitNickNameInput(ctx context.Context, b *bot.Bot, chatID, userID int64, inpunChan chan string, PatternName string, productID int) {
	defer close(inpunChan)
	select {
	case value := <-inpunChan:

		InfoPay := handlers.CreatePayURL(value, productID)

		if InfoPay.Response.URL != "" {
			kb := &models.InlineKeyboardMarkup{
				InlineKeyboard: [][]models.InlineKeyboardButton{
					{
						{Text: "Оплатить", URL: InfoPay.Response.URL},
						{Text: "Меню", CallbackData: "menu"},
					}, {},
				},
			}

			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:      chatID,
				Text:        "*" + bot.EscapeMarkdown("Оплата") + "*\n" + bot.EscapeMarkdown("После оплаты вы получите услугу на ник: ") + "*" + bot.EscapeMarkdown(value) + "*" + bot.EscapeMarkdown("\nК оплате: "+config.PriceProduct[PatternName]+"₽"),
				ReplyMarkup: kb,
				ParseMode:   models.ParseModeMarkdown,
			})
		} else {
			kb := &models.InlineKeyboardMarkup{
				InlineKeyboard: [][]models.InlineKeyboardButton{
					{
						{Text: "Повторить попытку: " + config.PriceProduct[PatternName] + "₽", CallbackData: PatternName},
						{Text: "Меню", CallbackData: "menu"},
					}, {},
				},
			}

			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:      chatID,
				Text:        "*" + bot.EscapeMarkdown("Ошибка") + "*" + bot.EscapeMarkdown("\nК сожалению, в данный момент оплата недоступна из-за технических неполадок. Мы уже работаем над устранением проблемы и скоро всё наладится. Пожалуйста, попробуйте позднее."),
				ParseMode:   models.ParseModeMarkdown,
				ReplyMarkup: kb,
			})
		}
	case <-time.After(30 * time.Second):
		kb := &models.InlineKeyboardMarkup{
			InlineKeyboard: [][]models.InlineKeyboardButton{
				{
					{Text: "Повторить попытку: " + config.PriceProduct[PatternName] + "₽", CallbackData: PatternName},
					{Text: "Меню", CallbackData: "menu"},
				}, {},
			},
		}

		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      chatID,
			Text:        "Время ожидания ввода истекло.",
			ReplyMarkup: kb,
		})
		if err != nil {
			log.Println("Ошибка отправки сообщения: ", err)
		}
		storage.UserInfoChat.Delete(userID) // Удаляем состояние.
	}
}
