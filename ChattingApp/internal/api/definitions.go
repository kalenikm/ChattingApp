package api

import "time"

type NewChatRequest struct {
	Title string `json:"title"`
}

type Chat struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type NewMessageRequest struct {
	Time time.Time `json:"time"`
	Name string    `json:"name"`
	Text string    `json:"text"`
}

type Message struct {
	Id     string    `json:"id"`
	ChatId string    `json:"chatId"`
	Time   time.Time `json:"time"`
	Name   string    `json:"name"`
	Text   string    `json:"text"`
}
