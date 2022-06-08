package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionString string = "mongodb://127.0.0.1:27017"

func GetMessagesByChatId(chatId string) (*[]Message, error) {
	client, ctx, cancel, err := connectToDb()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cancel()

	chattingAppDb := client.Database("chattingApp")
	messagesCollection := chattingAppDb.Collection("messages")

	chatObjectId := stringToObjectId(chatId)
	messagesCursor, err := messagesCollection.Find(ctx, bson.M{"chatId": chatObjectId})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer messagesCursor.Close(ctx)

	var messages []Message
	err = messagesCursor.All(ctx, &messages)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &messages, nil
}

func AddMessage(chatId string, message *Message) (primitive.ObjectID, error) {
	client, ctx, cancel, err := connectToDb()
	if err != nil {
		log.Println(err)
		return primitive.NilObjectID, err
	}
	defer cancel()

	chattingAppDb := client.Database("chattingApp")
	chatsCollection := chattingAppDb.Collection("chats")
	messagesCollection := chattingAppDb.Collection("messages")

	chatObjectId := stringToObjectId(chatId)
	var chat Chat
	if err = chatsCollection.FindOne(ctx, bson.M{"_id": chatObjectId}).Decode(&chat); err != nil {
		log.Println(err)
		return primitive.NilObjectID, fmt.Errorf("Chat with id(%v) does not exist", chatId)
	}

	result, err := messagesCollection.InsertOne(ctx, *message)
	if err != nil {
		log.Println(err)
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func GetChats() (*[]Chat, error) {
	client, ctx, cancel, err := connectToDb()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cancel()

	chattingAppDb := client.Database("chattingApp")
	chatsCollection := chattingAppDb.Collection("chats")

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

func AddChat(chat *Chat) (primitive.ObjectID, error) {
	client, ctx, cancel, err := connectToDb()
	if err != nil {
		log.Println(err)
		return primitive.NilObjectID, err
	}
	defer cancel()

	chattingAppDb := client.Database("chattingApp")
	chatsCollection := chattingAppDb.Collection("chats")

	result, err := chatsCollection.InsertOne(ctx, *chat)
	if err != nil {
		log.Println(err)
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

func connectToDb() (mongo.Client, context.Context, context.CancelFunc, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Println(err)
		return mongo.Client{}, nil, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Println(err)
		cancel()
		return mongo.Client{}, nil, nil, err
	}

	return *client, ctx, cancel, nil
}
