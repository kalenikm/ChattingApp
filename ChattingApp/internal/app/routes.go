package app

import (
	"github.com/labstack/echo/v4"
)

func (server *Server) Routes() *echo.Echo {
	e := server.echo

	v1 := e.Group("v1/api")

	// healt check
	v1.GET("/status", server.ApiStatus())

	v1.GET("/chats", server.GetChats())
	v1.POST("/chats", server.AddChat())

	v1.GET("/chats/:chatId/messages", server.GetMessages())
	v1.POST("/chats/:chatId/messages", server.AddMessage())

	v1.GET("/ws", server.WebSocket())

	return e
}
