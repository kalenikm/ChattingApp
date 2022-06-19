package app

import (
	"log"
	"net/http"

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
