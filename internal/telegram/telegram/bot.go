package telegram

import (
    "github.com/cvenkman/telegramBot/pkg/storage"
	tgbotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot *tgbotAPI.BotAPI
    tokenStorage storage.ITokenStorage
}

func NewBot(b *tgbotAPI.BotAPI, tr storage.ITokenStorage) *Bot {
	return &Bot{bot: b, tokenStorage: tr}
}

func (b *Bot) Run() error {
    updateConfig := tgbotAPI.NewUpdate(0)
    updateConfig.Timeout = 60
    updates := b.bot.GetUpdatesChan(updateConfig)


    for update := range updates {
        if update.Message == nil {
            continue
        }
        if update.Message.IsCommand() {
			b.HandleCommand(update.Message)
        } else {
			b.HandleMessage(update.Message)
		}
    }
	return nil
}
