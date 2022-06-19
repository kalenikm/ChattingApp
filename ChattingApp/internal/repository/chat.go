package repository

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatRepository interface {
	GetChats() (*[]Chat, error)
	AddChat(chat *Chat) (primitive.ObjectID, error)
}

func (s *storage) GetChats() (*[]Chat, error) {
	client := s.client
	ctx := s.context

	chattingAppDb := client.Database(dbName)
	chatsCollection := chattingAppDb.Collection(chatsCollectionName)

	chatsCursor, err := chatsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer chatsCursor.Close(ctx)

	var chats []Chat
	err = chatsCursor.All(ctx, &chats)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &chats, nil
}

func (s *storage) AddChat(chat *Chat) (primitive.ObjectID, error) {
	client := s.client
	ctx := s.context

	chattingAppDb := client.Database(dbName)
	chatsCollection := chattingAppDb.Collection(chatsCollectionName)

	result, err := chatsCollection.InsertOne(ctx, chat)
	if err != nil {
		log.Println(err)
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}
