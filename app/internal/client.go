package handlers

import (
	"github.com/gorilla/websocket"
	chat "chat/app/internal/handlers"
)


type ClientList map[*Client]bool


type Client struct {
	connection *websocket.Conn
	manager *chat.Manager
}


func NewClient(conn *websocket.Conn, manager *chat.Manager) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
	}
}