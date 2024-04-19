package main

import (
	// "chat/app/database"
	handler "chat/app/internal/handlers"
	"log"
	"net/http"
)


func main() {
	// database.ConnectDB()

	http.HandleFunc("/chat", handler.ChatTemplate)
	http.HandleFunc("/ws/chat", handler.ChatConnection)
	log.Println("http server started on :8000")

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
