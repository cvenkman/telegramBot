package telegram

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/cvenkman/telegramBot/pkg/models"
	"github.com/cvenkman/telegramBot/pkg/storage"
	// "github.com/cvenkman/telegramBot/pkg/storage"
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
		msg.Text = b.HandlePhotoCommands(message.Command())
	default:
		msg.Text = "I don't know that command"
	}
	if _, err := b.bot.Send(msg); err != nil {
		log.Println(err)
	}
}

func (b *Bot) HandlePhotoCommands(command string) string {
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

func (b *Bot) HandleMessage(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	if strings.Contains(message.Text, "пидор") || strings.Contains(message.Text, "пизда") ||
		strings.Contains(message.Text, "гей") {
		msg.Text = "сам такой"
	} else if strings.Contains(message.Text, "спасибо") {
		msg.Text = "обращайся)"
	} else if (strings.Contains(message.Text, "Link")) {
		err := b.tokenStorage.Save(message.Chat.ID, message.Text, storage.AccessTokens)
		if err != nil {
			log.Println(err)
		}
		msg.Text = "saved"
	} else {
		switch message.Text {
		case "get":
			msg.Text, _ = b.tokenStorage.Get(message.Chat.ID, storage.AccessTokens)
		case "да":
			msg.Text = "пизда"
		case "cat", "дай кота", "котя", "Cat", "Котя":
			msg.Text = b.HandlePhotoCommands("cat")
		case "keyboard", "Keyboard":
			msg.ReplyMarkup = numericKeyboard
		case "close", "Close":
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		case "dog", "Dog", "дай собаку":
			msg.Text = b.HandlePhotoCommands("dog")
		case "fox", "Fox", "лиса":
			msg.Text = b.HandlePhotoCommands("fox")
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
