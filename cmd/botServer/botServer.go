package main

import (
	"log"
	"os"
	tgbotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/cvenkman/telegramBot/internal/telegram/telegram"
	"github.com/boltdb/bolt"
	"github.com/cvenkman/telegramBot/pkg/storage"
	"github.com/cvenkman/telegramBot/pkg/storage/boltDB"
	// "flag"
	// "github.com/BurntSushi/toml"
)

var configPath string

// func init() {
// 	flag.StringVar(&configPath, "config-path", "configs/config.toml", "path to config file")
// }

func main() {
	/* set up tgbotAPI */
    bot, err := tgbotAPI.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
    if err != nil {
        log.Fatal("fatal error: telegram api token")
    }
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	/* create a database */
	db, err := bolt.Open("bot.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	tokenStorage := boltDB.NewTokenStorage(db)

	/* checking that buckets are created */
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(storage.AccessTokens))
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists([]byte(storage.RequestTokens))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	/* create a bot */
	myBot := telegram.NewBot(bot, tokenStorage)

	/* run bot */
	err = myBot.Run()
	if err != nil {
		log.Fatal(err)
	}
}
