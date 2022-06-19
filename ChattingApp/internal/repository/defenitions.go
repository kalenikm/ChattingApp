package repository

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var dbName = "chattingApp"
var chatsCollectionName = "chats"

//var messagesCollectionName = "messages"

type Message struct {
	Id     primitive.ObjectID `bson:"_id,omitempty"`
	ChatId primitive.ObjectID `bson:"chatId,omitempty"`
	Time   time.Time          `bson:"time,omitempty"`
	Name   string             `bson:"name,omitempty"`
	Text   string             `bson:"text,omitempty"`
}

type Chat struct {
	Id    primitive.ObjectID `bson:"_id,omitempty"`
	Title string             `bson:"title,omitempty"`
}
