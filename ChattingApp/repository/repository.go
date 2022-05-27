package repository

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionString string = "mongodb://127.0.0.1:27017"

func GetMessagesByChatId(chatId int) (*[]Message, error) {
	client, ctx, cancel, err := connectToDb()
	defer cancel()
	if err != nil {
		log.Println(err)
		return &[]Message{}, err
	}

	chattingAppDb := client.Database("chattingApp")
	messagesCollection := chattingAppDb.Collection("messages")

	messagesCursor, err := messagesCollection.Find(ctx, bson.M{"chatId": chatId})
	if err != nil {
		log.Println(err)
		return &[]Message{}, err
	}
	defer messagesCursor.Close(ctx)

	var messages []Message
	err = messagesCursor.All(ctx, &messages)
	if err != nil {
		return &[]Message{}, err
	}

	return &messages, nil
}

func AddMessage(message *Message) (primitive.ObjectID, error) {
	client, ctx, cancel, err := connectToDb()
	defer cancel()
	if err != nil {
		log.Println(err)
		return primitive.ObjectID{}, err
	}

	chattingAppDb := client.Database("chattingApp")
	messagesCollection := chattingAppDb.Collection("messages")

	result, err := messagesCollection.InsertOne(ctx, *message)
	if err != nil {
		log.Println(err)
		return primitive.ObjectID{}, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
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
