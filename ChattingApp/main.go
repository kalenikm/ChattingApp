package main

import (
	"chattingApp/repository"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/chats", func(c echo.Context) error {
		chats, err := repository.GetChats()
		if err != nil {
			log.Println(err)
			return err
		}

		return c.JSON(http.StatusOK, chats)
	})

	e.GET("/chats/:id", func(c echo.Context) error {
		chatId := c.Param("id")

		_, err := repository.GetChatById(chatId)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, err.Error())
		}

		messages, err := repository.GetMessagesByChatId(chatId)
		if err != nil {
			log.Println(err)
			return err
		}

		return c.JSON(http.StatusOK, messages)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
