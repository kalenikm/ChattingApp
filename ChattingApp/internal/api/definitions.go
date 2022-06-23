package api

import (
	"time"

	"github.com/gorilla/websocket"
)

type NewChatRequest struct {
	Title string `json:"title"`
}

type Chat struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type NewMessageRequest struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

type Message struct {
	Id     string    `json:"id"`
	ChatId string    `json:"chatId"`
	Time   time.Time `json:"time"`
	Name   string    `json:"name"`
	Text   string    `json:"text"`
}

type SocketPayload struct {
	Message string
}

type SocketResponse struct {
	Name    string
	Message string
}

type WebSocketConnection struct {
	Conn   *websocket.Conn
	ChatId string
	Name   string
}
