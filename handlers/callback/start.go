package callback

import (
	"context"
	"fmt"
	"telegarm/config"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func StartHandeler(ctx context.Context, b *bot.Bot, update *models.Update) {
	fmt.Println(update.Message)
	if update.Message != nil {
		kb := &models.InlineKeyboardMarkup{
			InlineKeyboard: [][]models.InlineKeyboardButton{
				{
					{Text: "Проходка: " + config.PriceProduct["buy_ticket"] + "₽", CallbackData: "info_ticket"},
					{Text: "Подписка Plus+: " + config.PriceProduct["buy_sponsor"] + "₽", CallbackData: "info_sponsor"},
				}, {
					{Text: "Помощь/Поддержка", CallbackData: "other"},
				},
			},
		}

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "*" + bot.EscapeMarkdown("Магазин") + "*" + "\n*" + bot.EscapeMarkdown("Почему тут донат?") + "*\nБольшое спасибо вам за материальную поддержку, мы не кладём всё в карман, а занимаемся улучшениями сервера и его дочерних ресурсов" + "\n*" + bot.EscapeMarkdown("\nНапоминание") + "*" + bot.EscapeMarkdown("\nВы можете не покупать проходку и попасть бесплатно во время бесплатного набора \n\nДенежные средтсва возврату не подлежат. Это является добровольным пожертвованием на развитие Сервера за которые вы получаете услугу в виде спасибо ") + "\n\n*" + bot.EscapeMarkdown("Как получить товар?") + "*" + bot.EscapeMarkdown("\nТовар вы получаете мгновенно на сервере по вашему нику который вы впишите при оплате."),
			ParseMode:   models.ParseModeMarkdown,
			ReplyMarkup: kb,
		})
	} else {
		fmt.Println("Received update without a message or chat.") // Example logging
	}
}
