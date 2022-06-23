package api

import (
	"log"

	"github.com/gorilla/websocket"
)

type WebSocketService interface {
	HandleConnection(currentConn *WebSocketConnection)
}

type webSocketService struct {
	messageService MessageService
	connections    []*WebSocketConnection
}

func NewWebSocketService(messageService MessageService) WebSocketService {
	return &webSocketService{
		messageService: messageService,
		connections:    make([]*WebSocketConnection, 0),
	}
}

func (s *webSocketService) HandleConnection(currentConn *WebSocketConnection) {
	s.connections = append(s.connections, currentConn)

	for {
		payload := SocketPayload{}
		err := currentConn.Conn.ReadJSON(&payload)
		if err != nil {
			if websocket.IsCloseError(err) {
				s.ejectConnection(currentConn)
				return
			}

			log.Println(err)
			continue
		}

		messageRequest := &NewMessageRequest{
			Name: currentConn.Name,
			Text: payload.Message,
		}

		_, err = s.messageService.AddMessage(currentConn.ChatId, messageRequest)
		if err != nil {
			log.Println(err)
			continue
		}

		s.broadcastMessage(currentConn, payload.Message)
	}
}

func (s *webSocketService) ejectConnection(currentConn *WebSocketConnection) {
	filtered := make([]*WebSocketConnection, 0)

	for _, conn := range s.connections {
		if conn != currentConn {
			filtered = append(filtered, conn)
		}
	}

	s.connections = filtered
}

func (s *webSocketService) broadcastMessage(currentConn *WebSocketConnection, message string) {
	for _, conn := range s.connections {
		if conn == currentConn || conn.ChatId != currentConn.ChatId {
			continue
		}

		response := SocketResponse{
			Name:    currentConn.Name,
			Message: message,
		}

		conn.Conn.WriteJSON(response)
	}
}
