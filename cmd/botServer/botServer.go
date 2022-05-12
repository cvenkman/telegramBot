package main

import (
	"log"
	"os"
	tgbotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/cvenkman/telegramBot/internal/telegram/telegram"
	"github.com/boltdb/bolt"
	"github.com/cvenkman/telegramBot/pkg/storage"
	"github.com/cvenkman/telegramBot/pkg/storage/boltDB"
)

/*
	TODO
	разделить main на функции
*/

func main() {
	/* настройка tgbotAPI */
    bot, err := tgbotAPI.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
    if err != nil {
        log.Fatal(err)
    }
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	/* создаем базу данных */
	db, err := bolt.Open("bot.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	tokenStorage := boltDB.NewTokenStorage(db)

	/* проверка что бакеты созданы */
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

	/* создаем бота */
	myBot := telegram.NewBot(bot, tokenStorage)

	/* запускаем бота */
	err = myBot.Run()
	if err != nil {
		log.Fatal(err)
	}
}




















// func main1() {
// 	botToken := "5389287107:AAEPaAIvxgV7fPmJLeAtMp4GvHQdLOgDdfM"
// 	// botAPI := "https://api.telegram.org/bot"
// 	// botURL := botAPI + botToken


// 	// подключаемся к боту с помощью токена
// 	bot, err := tgbotapi.NewBotAPI(botToken)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	bot.Debug = true
// 	log.Printf("Authorized on account %s", bot.Self.UserName)

// 	// инициализируем канал, куда будут прилетать обновления от API
// 	u := tgbotapi.NewUpdate(0)
// 	u.Timeout = 60

// 	updates, err := bot.GetUpdatesChan(u)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// читаем обновления из канала
// 	for update := range updates {
// 		switch {
// 		case update.Message != nil: // Если было прислано сообщение, то обрабатываем, так как могут приходить не только сообщения.
// 		OnMessage(bot, update.Message)
// 		}
// 	}
// }

// func OnMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
// 	// Пользователь, который написал боту
// 	userName := message.From.UserName

// 	// ID чата/диалога.
// 	// Может быть идентификатором как чата с пользователем
// 	// (тогда он равен UserID) так и публичного чата/канала
// 	chatID := message.Chat.ID

// 	log.Printf("[%s] %d", userName, chatID)

// 	spew.Dump(message) // выводим то что пришло (Для отладки!!!)

// 	var msg tgbotapi.Chattable

// 	switch {
// 	case message.Text != "": // Текстовое ли сообщение?
// 		msg = tgbotapi.NewMessage(chatID, message.Text)

// 	case message.Photo != nil: // Это фото?
		
// 		photoArray := *message.Photo
// 		photoLastIndex := len(photoArray) - 1
// 		photo := photoArray[photoLastIndex] // Получаем последний элемент массива (самую большую картинку)
// 		msg = tgbotapi.NewPhotoShare(chatID, photo.FileID)

// 	default:                                                 // Если не одно условие не сработало
// 		msg = tgbotapi.NewMessage(chatID, "Не реализовано") // Отправляется на тот тип сообщения, который ещё не реализован выше ^
// 	}

// 	// и отправляем его
// 	_, err := bot.Send(msg)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }
















// func runBot() {
// 	botToken := "5389287107:AAEPaAIvxgV7fPmJLeAtMp4GvHQdLOgDdfM"
// 	botAPI := "https://api.telegram.org/bot"
// 	botURL := botAPI + botToken
// 	offset := 0

// 	log.Println("start...")

// 	for {
// 		updates, err := getUpdates(botURL, offset)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		for _, update := range updates {
// 			err = sendMessage(botURL, update)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			offset = update.UpdateID + 1
// 		}
// 	}
// }

// func getUpdates(url string, offset int) ([]models.Update, error) {
// 	response, err :=  http.Get(url + "/getUpdates?offset=" + strconv.Itoa(offset))
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer response.Body.Close()

// 	body, err := ioutil.ReadAll(response.Body) // FIXME deprecated
// 	if err != nil {
// 		return nil, err
// 	}

// 	var updatesResponse models.Response

// 	err = json.Unmarshal(body, &updatesResponse)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return updatesResponse.Updates, nil
// }

// func sendMessage(url string, update models.Update) error {

// 	var botMessage models.BotMessage
// 	botMessage.ChatID = update.Message.Chat.ID
// 	botMessage.Text = "https://www.google.com/search?q=cat&sxsrf=ALiCzsadZDlUREXztkTARsBzzhvqE0fhew:1651690786338&tbm=isch&source=iu&ictx=1&vet=1&fir=U6GpUslQrdCfcM%252C_Xck-vPrSARjYM%252C_%253BhMjxy8pUhhc4QM%252C3aBlXpmFZqFG2M%252C_%253Boc0yOiQ9sK4GZM%252CfUaFTPQ4f7pyiM%252C_%253BY-AJkCjvG49YEM%252CUT0uQoxCDvb8yM%252C_%253BVRmao7FmvjTB6M%252CHkevFQZ5DYu7oM%252C_%253BAJAijU-By5qnBM%252Cw0PcXTTueMJSMM%252C_%253BBuyI8dg85sSx5M%252CvncSeXMVok58VM%252C_%253BRnJ5AJeGvMYVMM%252CfA78NAeh2cpfUM%252C_%253BztjznUfJl4TsCM%252CmW4cOYEIUPJVrM%252C_%253ByY42gJ5WMC1QdM%252CwYHY6iexestFEM%252C_&usg=AI4_-kQ6HQCCbsB73TNaVhQhpmuK9UtmiA&sa=X&ved=2ahUKEwiKwOTUw8b3AhWCmIsKHeFTBJkQ9QF6BAgFEAE#imgrc=BuyI8dg85sSx5M&imgdii=_bX9dbChNXSEFM"

// 	buf, err := json.Marshal(&botMessage)
// 	if err != nil {
// 		return err
// 	}



// 	// buf := new(bytes.Buffer)

// 	// m := image.NewRGBA(image.Rect(0, 0, 16, 16))
// 	// clr := color.RGBA{B: 0, A: 0}
// 	// draw.Draw(m, m.Bounds(), &image.Uniform{C: clr}, image.ZP, draw.Src)

// 	// var img image.Image = m
// 	// if err := jpeg.Encode(buf, img, nil); err != nil {
// 	// 	return err
// 	// }
// 	//buf.Bytes()


// 	// _, err = http.Post(url + "/sendMessage", "application/json", bytes.NewBuffer(buf))
// 	_, err = http.Post(url + "/sendPhoto", "application/json", bytes.NewBuffer(buf))
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// //https://lh3.googleusercontent.com/YGDk-aWozj7Vv-QG0sVk6vLXAiOGBnCuPiQYmOHd0NBOlLMdpS6GkrTBHztUDbNISSoPQcLtYzwHHCUiDw8JCbDaU9r73oxW3Tfc7g=s0