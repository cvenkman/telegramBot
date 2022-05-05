package telegram

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

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
	case "cat":
		str, err := getImageURL()
		if err != nil {
			log.Println(err)
		}
		msg.Text = str
	case "start":
		msg.Text = "ты гей?"
		
	// case "kitty":
	// buf := new(bytes.Buffer)

	// m := image.NewRGBA(image.Rect(0, 0, 16, 16))
	// clr := color.RGBA{B: 0, A: 0}
	// draw.Draw(m, m.Bounds(), &image.Uniform{C: clr}, image.ZP, draw.Src)

	// var img image.Image = m
	// if err := jpeg.Encode(buf, img, nil); err != nil {
	// 	log.Println(err)
	// }

	// // var reader io.Reader
 
	// // file := tgbotapi.FileBytes{
	// // 	Name: "image.jpg",
	// // 	Bytes: buf.Bytes(),
	// // }

	// fmt.Println(update.Message.Text)
	// // msg.Text = file.SendData()
	// msg.Text = string(buf.Bytes())
	// if _, err := bot.Send(msg); err != nil {
	// 	log.Println(err)
	// }
	default:
		msg.Text = "I don't know that command"
	}
	if _, err := b.Bot.Send(msg); err != nil {
		log.Println(err)
	}
}

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/cat"),
	),
)

// func (b *Bot) HandleHiMessage(message *tgbotapi.Message) {
// 	msg := tgbotapi.NewMessage(message.Chat.ID, "Привет")
// 	if _, err := b.Bot.Send(msg); err != nil {
// 		log.Println(err)
// 	}
// }


/* разделить на разные Handle */
func (b *Bot) HandleMessage(message *tgbotapi.Message) {
	var err error
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	if strings.Contains(message.Text, "пидор") || strings.Contains(message.Text, "пизда") ||
		strings.Contains(message.Text, "гей") {
		msg.Text = "сам такой"
	} else if strings.Contains(message.Text, "спасибо") {
		msg.Text = "обращайся)"
	} else {
		switch message.Text {
		case "да":
			msg.Text = "пизда"
		case "дай кота":
			msg.Text, err = getImageURL()
			if err != nil {
				log.Println(err)
			}
		case "keyboard", "Keyboard":
			msg.ReplyMarkup = numericKeyboard
		case "close":
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		default:
			msg.Text = "Ну и что это?"
		}
	}

	if _, err := b.Bot.Send(msg); err != nil {
		log.Println(err)
	}
}

type AutoGenerated []struct {
	URL string `json:"url"`
}

func getImageURL() (string, error) {
	response, err := http.Get("https://api.thecatapi.com/v1/images/search")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body) // FIXME deprecated
	if err != nil {
		return "", err
	}

	var updatesResponse AutoGenerated

	err = json.Unmarshal(body, &updatesResponse)
	if err != nil {
		return "", err
	}

	return updatesResponse[0].URL, nil
}