package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"fmt"
)

type Bot struct {
	Bot *tgbotapi.BotAPI
}

func NewBot(b *tgbotapi.BotAPI) *Bot {
	return &Bot{Bot: b}
}

func (b *Bot) Run() error {
    // b.Bot, err = tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
    // if err != nil {
    //     log.Fatal(err)
    // }

    updateConfig := tgbotapi.NewUpdate(0)
    updateConfig.Timeout = 60
    updates := b.Bot.GetUpdatesChan(updateConfig)

	

    for update := range updates {
        if update.Message == nil {
            continue
        }
        if update.Message.IsCommand() {
			b.HandleCommand(update.Message)
        } else {
			b.HandleMessage(update.Message)
		}
		fmt.Println(update.Message.Text)
    }
	return nil
}
