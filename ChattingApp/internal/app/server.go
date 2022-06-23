package app

import (
	"chattingApp/internal/api"
	"log"

	"github.com/labstack/echo/v4"
)

type Server struct {
	echo             *echo.Echo
	chatService      api.ChatService
	messageService   api.MessageService
	webSocketService api.WebSocketService
}

func NewServer(echo *echo.Echo, chatService api.ChatService, messageService api.MessageService, webSocketService api.WebSocketService) *Server {
	return &Server{
		echo:             echo,
		chatService:      chatService,
		messageService:   messageService,
		webSocketService: webSocketService,
	}
}

func (server *Server) Run() error {
	e := server.Routes()

	err := e.Start(":1323")

	if err != nil {
		log.Fatal("Server start failed")
	}

	return nil
}
