package youtube

type YoutubeResponse struct {
	Items []Item	`json:"items"`
}

type Item struct {
	ID IDInfo	`json:"id"`
}

type IDInfo struct {
	VideoId string	`json:"videoId"`
}

/* json response example from youtube api
{
	"items": [
		{
			"id": {
			"videoId": "6uddGul0oAc"
			}
		},
		{
			"id": {
			"videoId": "lVpmZnRIMKs"
			}
		}
	]
}
*/