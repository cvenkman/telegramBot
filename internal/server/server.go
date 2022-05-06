package server

// import (
// 	"log"
// 	"net/http"
// )

// func Run() error {
// 	http.HandleFunc("/", imgHandler)
// 	http.HandleFunc("/favicon", faviconHandler)
// 	http.HandleFunc("/ping", pingHandler)
// 	http.HandleFunc("/robots.txt", robotsHandler)

// 	log.Println("start server...")
// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func imgHandler(w http.ResponseWriter, r *http.Request) {
// 	write(w, "img")
// }

// func faviconHandler(w http.ResponseWriter, r *http.Request) {
// 	write(w, "favicon")
// }

// func pingHandler(w http.ResponseWriter, r *http.Request) {
// 	write(w, "pong")
// }

// func robotsHandler(w http.ResponseWriter, r *http.Request) {
// 	write(w, "robots")
// }

// func write(w http.ResponseWriter, msg string) {
// 	_, err := w.Write([]byte([]byte(msg)))
// 	if err != nil {
// 		log.Println(err)
// 	}
// }