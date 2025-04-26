package callback

import (
	"context"
	"telegarm/config"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})

	if update.CallbackQuery.Data == "menu" {
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

		b.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
			MessageID:   update.CallbackQuery.Message.Message.ID,
			Text:        "*" + bot.EscapeMarkdown("Магазин") + "*" + "\n*" + bot.EscapeMarkdown("Почему тут донат?") + "*\nБольшое спасибо вам за материальную поддержку, мы не кладём всё в карман, а занимаемся улучшениями сервера и его дочерних ресурсов" + "\n*" + bot.EscapeMarkdown("\nНапоминание") + "*" + bot.EscapeMarkdown("\nВы можете не покупать проходку и попасть бесплатно во время бесплатного набора \n\nДенежные средтсва возврату не подлежат. Это является добровольным пожертвованием на развитие Сервера за которые вы получаете услугу в виде спасибо ") + "\n\n*" + bot.EscapeMarkdown("Как получить товар?") + "*" + bot.EscapeMarkdown("\nТовар вы получаете мгновенно на сервере по вашему нику который вы впишите при оплате."),
			ParseMode:   models.ParseModeMarkdown,
			ReplyMarkup: kb,
		})
	}

	if update.CallbackQuery.Data == "other" {
		kb := &models.InlineKeyboardMarkup{
			InlineKeyboard: [][]models.InlineKeyboardButton{
				{
					{Text: "FAQ", URL: "https://wiki.netlos.ru/"},
				}, {
					{Text: "Политика обработки персональных данных", URL: "https://1drv.ms/w/c/f8d83ab7cc7b02ef/EdmnmJh194RKkEDCEDajwjwBipxq1NeEzf8iebMXhuEr-A?e=AyV445"},
				}, {
					{Text: "Меню", CallbackData: "menu"},
				},
			},
		}

		b.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
			MessageID:   update.CallbackQuery.Message.Message.ID,
			Text:        "*" + bot.EscapeMarkdown("Инструкция по оплате банковскими картами") + "*" + bot.EscapeMarkdown("\n"+`К оплате принимаются платежные карты: Visa, MasterCard, Mir и UnionPay. Для оплаты товара банковской картой при оформлении заказа в магазине напишите свой ник и нажмите на кнопку "оплатить". При оплате заказа, обработка платежа происходит на авторизационной странице банка, где Вам необходимо ввести данные Вашей банковской карты: тип карты, номер карты, срок действия карты, CVC2/CVV2 код.`) + "\n\n*" + bot.EscapeMarkdown("Поддержка") + "*" + bot.EscapeMarkdown("\nDiscord: https://discord.gg/5Rkn5ecjq2"),
			ParseMode:   models.ParseModeMarkdown,
			ReplyMarkup: kb,
		})
	}

	if update.CallbackQuery.Data == "info_ticket" {
		kb := &models.InlineKeyboardMarkup{
			InlineKeyboard: [][]models.InlineKeyboardButton{
				{
					{Text: "Купить за " + config.PriceProduct["buy_ticket"] + "₽", CallbackData: "buy_ticket"},
					{Text: "Меню", CallbackData: "menu"},
				}, {},
			},
		}

		b.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
			MessageID:   update.CallbackQuery.Message.Message.ID,
			Text:        "*" + bot.EscapeMarkdown("Проходка: "+config.PriceProduct["buy_ticket"]+"₽") + "*" + bot.EscapeMarkdown("\nКупив вы получается доступ к серверу.") + "*" + bot.EscapeMarkdown("\n\nПодробнее: ") + "*" + "\nВо время нового сезона или межсезонья вас не будут проверять на правила и сразу же пустят на сервер, даже если вы не будете играть несколько месяцев",
			ParseMode:   models.ParseModeMarkdown,
			ReplyMarkup: kb,
		})
	}
	if update.CallbackQuery.Data == "info_sponsor" {
		kb := &models.InlineKeyboardMarkup{
			InlineKeyboard: [][]models.InlineKeyboardButton{
				{
					{Text: "Купить за: " + config.PriceProduct["buy_sponsor"] + "₽", CallbackData: "buy_sponsor"},
					{Text: "Меню", CallbackData: "menu"},
				}, {},
			},
		}

		description := bot.EscapeMarkdown("Это серверная подписка на месяц") + "\n"
		description += bot.EscapeMarkdown("С помощью этой подписки вы можете:") + "\n"
		description += "• " + bot.EscapeMarkdown("Менять размер персонажа") + "\n"
		description += "• " + bot.EscapeMarkdown("Использование Со і (Проверять действие с блоком/объектом)") + "\n"
		description += "• " + bot.EscapeMarkdown("Выделяющая надпись Plus+ в дискорде и в игре") + "\n"
		description += "• " + bot.EscapeMarkdown("Информация с координатами вашей смерти после смерти") + "\n"
		description += "• " + bot.EscapeMarkdown("Можете использовать партиклы") + "\n"
		description += "• " + bot.EscapeMarkdown("Отключение фантомов") + "\n"
		description += "• " + bot.EscapeMarkdown("Создание пластинок со своей песней") + "\n\n"
		description += bot.EscapeMarkdown("КОМАНДЫ") + "\n"
		description += bot.EscapeMarkdown("/scale set <НИК> 0.6") + "\n"
		description += bot.EscapeMarkdown("<- Размер") + "\n"
		description += bot.EscapeMarkdown("/co i") + "\n"
		description += bot.EscapeMarkdown("<- Логи вкл/выкл") + "\n"
		description += bot.EscapeMarkdown("/pp") + "\n"
		description += bot.EscapeMarkdown("<- Красивые эффекты персонажей, ломание блоков, бег, установка блоков, полёт и тд") + "\n\n"
		description += bot.EscapeMarkdown("ПРАВИЛА") + "\n"
		description += "• " + bot.EscapeMarkdown("Минимальный размер персонажа 0.6 максимальный 1.3, если выдадите больше или меньше, то у вас отберут возможность использование команды.") + "\n"
		description += "• " + bot.EscapeMarkdown("Логи (СО І) прошу не абузить") + "\n"
		description += "• " + bot.EscapeMarkdown("Использовать команды для выгоды других запрещено") + "\n"
		description += "• " + bot.EscapeMarkdown("Выдавать размер другому человеку запрещено (отберут возможность)")

		b.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
			MessageID:   update.CallbackQuery.Message.Message.ID,
			Text:        "*" + bot.EscapeMarkdown("Подписка Plus+: "+config.PriceProduct["buy_sponsor"]+"₽ (МЕСЯЦ)") + "*\n" + description,
			ParseMode:   models.ParseModeMarkdown,
			ReplyMarkup: kb,
		})
	}
}
