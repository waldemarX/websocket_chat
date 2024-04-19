package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type PageData struct {
	Messages []string
}

func ChatTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("app/templates/chat.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := PageData{
		Messages: []string{"Hello", "World"},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


var upgrader = websocket.Upgrader{
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


type Manager struct {
}


func NewManager() *Manager {
	return &Manager{}
}


func ChatConnection(w http.ResponseWriter, r *http.Request) {
	log.Println("New connection from")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
}
