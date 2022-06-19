package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Storage interface {
	GetChats() (*[]Chat, error)
	AddChat(chat *Chat) (primitive.ObjectID, error)
}

type storage struct {
	client  *mongo.Client
	context context.Context
}

func NewStorage(client *mongo.Client, ctx context.Context) Storage {
	return &storage{
		client:  client,
		context: ctx,
	}
}
