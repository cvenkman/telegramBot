package telegram

import (
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
		msg.Text = `I understand /sayhi, /cat, /dog and /fox commands.
					You can also write animal (support only cat, dog and fox) to get a photo of this animal.

					One more thing I can do: save smth (write save [your information]) and then get it (write get),
					e.g., save hi, get.
					
					Then you can get the latest video from the YouTube channel (but url must contains channel id),
					e.g, video https://www.youtube.com/channel/UC5A-Wp9ujcr5g9sYagAafEA`

	case "sayhi":
		msg.Text = "Hi :)"
	case "start":
		msg.Text = "Hi! Write /help to see a list of commands"
	case "cat", "dog", "fox":
		msg.Text = b.handlePhotoCommands(message.Command())
	default:
		msg.Text = `I don't know that command
					/help`
	}
	if _, err := b.bot.Send(msg); err != nil {
		log.Println(err)
	}
}

func (b *Bot) handlePhotoCommands(command string) string {
	var msg string
	var err error

	switch command {
	case "cat":
		var d models.CatURL
		msg, err = getImageURL("https://api.thecatapi.com/v1/images/search", d)
	case "dog":
		var f models.DogURL
		msg, err = getImageURL("https://random.dog/woof.json", f)
	case "fox":
		var f models.FoxURL
		msg, err = getImageURL("https://randomfox.ca/floof/", f)
	}
	if err != nil {
		log.Println(err)
		msg = "server error: can't get image"
	}
	return msg
}

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/cat"),
		tgbotapi.NewKeyboardButton("/dog"),
		tgbotapi.NewKeyboardButton("/fox"),
	),
)

func (b *Bot) HandleMessage(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	var err error

	if strings.Contains(message.Text, "спасибо") {
		msg.Text = "обращайся)"
	} else if strings.Contains(message.Text, "Save") || strings.Contains(message.Text, "save"){
		err := b.tokenStorage.Save(message.Chat.ID, message.Text, storage.AccessTokens)
		if err != nil {
			log.Println(err)
			msg.Text = "server error: can't save"
		}
		msg.Text = "saved"
	} else if strings.Contains(message.Text, "Video") || strings.Contains(message.Text, "video") {
		channelURL := strings.Trim(message.Text, " \t")
		channelURL = channelURL[5:]

		msg.Text, err = youtube.GetLastVideo(channelURL)
		if err != nil {
			log.Println(err)
			msg.Text = "not valid URL or maybe my bad"
		}
	} else {
		switch message.Text {
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
			msg.Text = "I don't know this :( /help"
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

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return s.Unmarshal(body)
}
