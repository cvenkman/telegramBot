package models

import (
	"encoding/json"
	"fmt"
)

type ToUnmarshal interface {
	Unmarshal(data []byte) (string, error)
}

type CatURL []struct {
	URL string `json:"url"`
}

type DogURL struct {
	URL string `json:"url"`
}

type FoxURL struct {
	URL string `json:"image"`
}

func (c CatURL) Unmarshal(data []byte) (string, error) {
	var updatesResponse CatURL
	err := json.Unmarshal(data, &updatesResponse)
	if err != nil {
		return "", err
	}

	return updatesResponse[0].URL, nil
}

func (c DogURL) Unmarshal(data []byte) (string, error) {
	var updatesResponse DogURL
	err := json.Unmarshal(data, &updatesResponse)
	if err != nil {
		return "", err
	}

	return updatesResponse.URL, nil
}

func (c FoxURL) Unmarshal(data []byte) (string, error) {
	fmt.Println("--- ")
	var updatesResponse FoxURL
	err := json.Unmarshal(data, &updatesResponse)
	fmt.Println("--- ", updatesResponse.URL)
	if err != nil {
		return "", err
	}

	return updatesResponse.URL, nil
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
