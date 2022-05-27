package main

import (
	"chattingApp/repository"
	"fmt"
	"log"
	"time"
)

func main() {
	newMessage := &repository.Message{
		ChatId: 1,
		Time:   time.Now(),
		Name:   "Mikhail",
		Text:   "Hi there!",
	}

	objectId, err := repository.AddMessage(newMessage)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted object id: %v\n", objectId)

	messages, err := repository.GetMessagesByChatId(1)
	if err != nil {
		log.Fatal(err)
	}
	for _, value := range *messages {
		fmt.Printf("Id: %v; ChatId: %v; Time: %v, Name: %v, Text: %v;\n", value.Id, value.ChatId, value.Time, value.Name, value.Text)
	}
}
