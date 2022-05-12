package telegram

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/cvenkman/telegramBot/internal/telegram/models"
	"github.com/cvenkman/telegramBot/internal/youtube/youtube"
	"github.com/cvenkman/telegramBot/pkg/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) HandleCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	switch message.Command() {
	case "help":
		msg.Text = "I understand /sayhi, /status or /cat."
	case "sayhi":
		msg.Text = "Hi :)"
	case "status":
		msg.Text = "I'm ok."
	case "start":
		msg.Text = "ты гей?"
	case "cat", "dog", "fox":
		msg.Text = b.handlePhotoCommands(message.Command())
	default:
		msg.Text = "I don't know that command"
	}
	if _, err := b.bot.Send(msg); err != nil {
		log.Println(err)
	}
}

func (b *Bot) handlePhotoCommands(command string) string {
	var str string
	var err error

	switch command {
	case "cat":
		var d models.CatURL
		str, err = getImageURL("https://api.thecatapi.com/v1/images/search", d)
		if err != nil {
			log.Println(err)
		}
	case "dog":
		var f models.DogURL
		str, err = getImageURL("https://random.dog/woof.json", f)
		if err != nil {
			log.Println(err)
		}
	case "fox":
		var f models.FoxURL
		str, err = getImageURL("https://randomfox.ca/floof/", f)
		if err != nil {
			log.Println(err)
		}
	}
	return str
}

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/cat"),
		tgbotapi.NewKeyboardButton("/dog"),
		tgbotapi.NewKeyboardButton("/fox"),
	),
)

// func (b *Bot) HandleHiMessage(message *tgbotapi.Message) {
// 	msg := tgbotapi.NewMessage(message.Chat.ID, "Привет")
// 	if _, err := b.Bot.Send(msg); err != nil {
// 		log.Println(err)
// 	}
// }


//UC5A-Wp9ujcr5g9sYagAafEA
func (b *Bot) HandleMessage(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	var err error

	if strings.Contains(message.Text, "спасибо") {
		msg.Text = "обращайся)"
	} else if (strings.Contains(message.Text, "Link")) {
		err := b.tokenStorage.Save(message.Chat.ID, message.Text, storage.AccessTokens)
		if err != nil {
			log.Println(err)
		}
		msg.Text = "saved"
	} else {
		switch message.Text {
		case "video":
			msg.Text, err = youtube.GetLastVideo("https://www.youtube.com/channel/UCpOH8JsphMAVN7yTqJPQDew")
			fmt.Println("--------- ", msg.Text)
			if err != nil {
				fmt.Println("---------1 ", err)
				log.Println(err)
			}
		case "get":
			msg.Text, _ = b.tokenStorage.Get(message.Chat.ID, storage.AccessTokens)
		case "cat", "кот", "Cat", "Кот":
			msg.Text = b.handlePhotoCommands("cat")
		case "dog", "Dog", "Собака", "собака":
			msg.Text = b.handlePhotoCommands("dog")
		case "fox", "Fox", "лиса", "Лиса":
			msg.Text = b.handlePhotoCommands("fox")
		case "keyboard", "Keyboard":
			msg.ReplyMarkup = numericKeyboard
		case "close", "Close":
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		default:
			msg.Text = "Ну и что это?"
		}
	}

	if _, err := b.bot.Send(msg); err != nil {
		log.Println(err)
	}
}

func getImageURL(url string, s models.ToUnmarshal) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body) // FIXME deprecated
	if err != nil {
		return "", err
	}

	return s.Unmarshal(body)
}
