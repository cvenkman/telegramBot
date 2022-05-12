package youtube

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	youtubeModels "github.com/cvenkman/telegramBot/internal/youtube/models"
)

// GET https://youtube.googleapis.com/youtube/v3/search?part=id&channelId=UC5A-Wp9ujcr5g9sYagAafEA&maxResults=1&order=rating&key=[YOUR_API_KEY] HTTP/1.1

// Authorization: Bearer [YOUR_ACCESS_TOKEN]
// Accept: application/json



// FIXME убрать из методов channelURL или вызывать makeRequest из GetLastVideo


const SearchURL = "https://www.googleapis.com/youtube/v3"
const VideoURL = "https://www.youtube.com/watch?v="
const MaxResults = 1
// return url to last video from channelURL
func GetLastVideo(channelURL string) (string, error) {
	videos, err := getVideoItems(channelURL)

	if err != nil {
		log.Println() // FIXME
		return "", err
	}
	if len(videos) < 1 {
		return "", errors.New("video not found")
	}
	return VideoURL + videos[0].ID.VideoId, nil
}

func getVideoItems(channelURL string) ([]youtubeModels.Item, error) {
	request, err := makeRequest(channelURL)
	if err != nil {
		return nil, err
	}
	// клиент который будет обрабатывать запрос
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var restResponse youtubeModels.YoutubeResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Items, nil
}

func makeRequest(channelURL string) (*http.Request, error) {
	lastSlash := strings.LastIndex(channelURL, "/")
	channelID := channelURL[lastSlash + 1 :]
	// channelID := ("UC5A-Wp9ujcr5g9sYagAafEA")
	request, err := http.NewRequest("GET", SearchURL, nil)
	if err != nil {
		return nil, err
	}
 
	query := request.URL.Query()
	query.Add("part", "id")
	//UC5A-Wp9ujcr5g9sYagAafEA
	query.Add("channelId", channelID)
	query.Add("maxResults", strconv.Itoa(MaxResults))
	query.Add("order", "rating")
	query.Add("key", os.Getenv("YOUTUBE_ACCESS_TOKEN")) // FIXME ADD ERR CHECK

	request.URL.RawQuery = query.Encode()
	fmt.Println("+++++++++ ", request.URL.RawQuery)
	return request, nil
}
