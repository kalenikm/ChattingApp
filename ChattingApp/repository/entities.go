package repository

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	Id     primitive.ObjectID `bson:"_id,omitempty"`
	ChatId int                `bson:"chatId,omitempty"`
	Time   time.Time          `bson:"time,omitempty"`
	Name   string             `bson:"name,omitempty"`
	Text   string             `bson:"text,omitempty"`
}
