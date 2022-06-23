package app

import (
	"chattingApp/internal/api"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

func (s *Server) ApiStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Add("Content-Type", "application/json")

		response := map[string]string{
			"status": "success",
			"data":   "chattingApp API running smoothly",
		}

		return c.JSON(http.StatusOK, response)
	}
}

func (s *Server) GetChats() echo.HandlerFunc {
	return func(c echo.Context) error {
		chats, err := s.chatService.GetChats()
		if err != nil {
			log.Println(err)
			return err
		}

		return c.JSON(http.StatusOK, chats)
	}
}

func (s *Server) AddChat() echo.HandlerFunc {
	return func(c echo.Context) error {
		chatRequest := &api.NewChatRequest{}

		if err := c.Bind(chatRequest); err != nil {
			log.Println(err)
			return err
		}

		createdId, err := s.chatService.CreateChat(chatRequest)
		if err != nil {
			log.Println(err)
			return err
		}

		return c.JSON(http.StatusCreated, createdId)
	}
}

func (s *Server) GetMessages() echo.HandlerFunc {
	return func(c echo.Context) error {
		chatId := c.Param("chatId")
		messages, err := s.messageService.GetMessages(chatId)
		if err != nil {
			log.Println(err)
			return err
		}

		return c.JSON(http.StatusOK, messages)
	}
}

func (s *Server) AddMessage() echo.HandlerFunc {
	return func(c echo.Context) error {
		chatId := c.Param("chatId")

		messageRequest := &api.NewMessageRequest{}

		if err := c.Bind(messageRequest); err != nil {
			log.Println(err)
			return err
		}

		createdId, err := s.messageService.AddMessage(chatId, messageRequest)
		if err != nil {
			log.Println(err)
			return err
		}

		return c.JSON(http.StatusCreated, createdId)
	}
}

var upgrader = websocket.Upgrader{}

func (s *Server) WebSocket() echo.HandlerFunc {
	return func(c echo.Context) error {
		conn, err := upgrader.Upgrade(c.Response(), c.Request(), c.Response().Header())
		if err != nil {
			log.Println(err)
		}

		name := c.QueryParam("name")
		chatId := c.QueryParam("chatId")
		currentConn := api.WebSocketConnection{Conn: conn, ChatId: chatId, Name: name}

		go s.webSocketService.HandleConnection(&currentConn)

		return nil
	}
}
