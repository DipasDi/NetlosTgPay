package input

import (
	"context"
	"telegarm/storage"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func HandleNameInput(ctx context.Context, bot *bot.Bot, update *models.Update) {

	if update.Message == nil {
		return
	}

	value, ok := storage.UserInfoChat.Load(update.Message.From.ID)
	if !ok {
		return
	}

	inputChan := value.(chan string)
	inputChan <- update.Message.Text
	storage.UserInfoChat.Delete(update.Message.From.ID)
}
