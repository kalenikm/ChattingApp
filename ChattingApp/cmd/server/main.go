package main

import (
	"chattingApp/internal/api"
	"chattingApp/internal/app"
	"chattingApp/internal/repository"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\\n", err)
		os.Exit(1)
	}
}

func run() error {
	connectionString := "mongodb://127.0.0.1:27017"

	client, ctx, cancel, err := connectToDb(connectionString)
	if err != nil {
		return err
	}
	defer cancel()

	storage := repository.NewStorage(client, ctx)
	e := echo.New()

	chatService := api.NewChatService(storage)
	messageService := api.NewMessageService(storage)
	webSocketService := api.NewWebSocketService(messageService)

	server := app.NewServer(e, chatService, messageService, webSocketService)

	err = server.Run()

	if err != nil {
		return err
	}

	return nil
}

func connectToDb(connectionString string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Println(err)
		return &mongo.Client{}, nil, nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	err = client.Connect(ctx)
	if err != nil {
		log.Println(err)
		cancel()
		return &mongo.Client{}, nil, nil, err
	}

	return client, ctx, cancel, nil
}
