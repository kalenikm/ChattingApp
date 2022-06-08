package main

import (
	"chattingApp/repository"
	"fmt"
	"log"
	"time"
)

func main() {
	var chat repository.Chat

	chats, err := repository.GetChats()
	if err != nil {
		log.Fatal(err)
	}

	if len(*chats) > 0 {
		chat = (*chats)[0]
	} else {
		chat = repository.Chat{
			Title: "Another Chat",
		}

		createId, err := repository.AddChat(&chat)
		if err != nil {
			log.Fatal(err)
		}

		chat.Id = createId
	}

	newMessage := &repository.Message{
		ChatId: chat.Id,
		Time:   time.Now(),
		Name:   "Mikhail",
		Text:   "Hi there!",
	}

	objectId, err := repository.AddMessage(chat.Id.Hex(), newMessage)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted object id: %v\n", objectId)

	messages, err := repository.GetMessagesByChatId(chat.Id.Hex())
	if err != nil {
		log.Fatal(err)
	}
	for _, value := range *messages {
		fmt.Printf("Id: %v; ChatId: %v; Time: %v, Name: %v, Text: %v;\n", value.Id, value.ChatId, value.Time, value.Name, value.Text)
	}
}
