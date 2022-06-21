package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageRepository interface {
	GetMessages(chatId string) (*[]Message, error)
	AddMessage(chatId string, message *Message) (primitive.ObjectID, error)
}

func (s *storage) GetMessages(chatId string) (*[]Message, error) {
	client := s.client
	ctx := s.context

	chattingAppDb := client.Database(dbName)
	messagesCollection := chattingAppDb.Collection(messagesCollectionName)

	chatObjectId := stringToObjectId(chatId)
	messagesCursor, err := messagesCollection.Find(ctx, bson.M{"chatId": chatObjectId})
	if err != nil {
		return nil, err
	}
	defer messagesCursor.Close(ctx)

	var messages []Message
	err = messagesCursor.All(ctx, &messages)
	if err != nil {
		return nil, err
	}

	return &messages, nil
}

func (s *storage) AddMessage(chatId string, message *Message) (primitive.ObjectID, error) {
	client := s.client
	ctx := s.context

	chattingAppDb := client.Database(dbName)
	chatsCollection := chattingAppDb.Collection(chatsCollectionName)
	messagesCollection := chattingAppDb.Collection(messagesCollectionName)

	chatObjectId := stringToObjectId(chatId)
	var chat Chat
	if err := chatsCollection.FindOne(ctx, bson.M{"_id": chatObjectId}).Decode(&chat); err != nil {
		return primitive.NilObjectID, err
	}

	message.ChatId = chatObjectId
	result, err := messagesCollection.InsertOne(ctx, *message)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func stringToObjectId(id string) primitive.ObjectID {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID
	}

	return objId
}
