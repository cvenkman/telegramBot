package models

/***** get ******/

type Response struct {
	Ok bool			`json:"ok"`
	Updates []Update	`json:"result"`
}

type Update struct {
	UpdateID int	`json:"update_id"`
	Message Message `json:"message"`
}

type Message struct {
	Chat Chat		`json:"chat"`
	Text string		`json:"text"`
}

type Chat struct {
	ID int			`json:"id"`
	Username string	`json:"username"`
}


/***** send ******/

type BotMessage struct {
	ChatID int		`json:"chat_id"`
	Text string		`json:"text"`
}












/*
response json example

{"ok":true, "result":
	[{"update_id":323294379, "message":
		{"message_id":1,
			"from": {"id":859479894, "is_bot":false, "first_name":"\u0410\u0440\u0438\u043d\u0430", "username":"Gjkos", "language_code":"en"},
			"chat": {"id":859479894, "first_name":"\u0410\u0440\u0438\u043d\u0430", "username":"Gjkos", "type":"private"},
			"date":1651663415,
			"text":"/start",
			"entities": [{"offset":0,"length":6,"type":"bot_command"}]
		}
	}]
}
*/