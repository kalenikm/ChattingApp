package api

import (
	"chattingApp/internal/repository"
	"time"
)

type MessageService interface {
	GetMessages(chatId string) (*[]Message, error)
	AddMessage(chatId string, message *NewMessageRequest) (string, error)
}

type messageService struct {
	storage repository.MessageRepository
}

func NewMessageService(messageRepo repository.MessageRepository) MessageService {
	return &messageService{
		storage: messageRepo,
	}
}

func (s *messageService) GetMessages(chatId string) (*[]Message, error) {
	messageEntities, err := s.storage.GetMessages(chatId)
	if err != nil {
		return nil, err
	}

	messages := make([]Message, len(*messageEntities))
	for i, message := range *messageEntities {
		messages[i].Id = message.Id.Hex()
		messages[i].ChatId = message.ChatId.Hex()
		messages[i].Time = message.Time
		messages[i].Name = message.Name
		messages[i].Text = message.Text
	}

	return &messages, nil
}

func (s *messageService) AddMessage(chatId string, message *NewMessageRequest) (string, error) {
	messageEntity := &repository.Message{
		Time: time.Now(),
		Name: message.Name,
		Text: message.Text,
	}

	createdId, err := s.storage.AddMessage(chatId, messageEntity)
	if err != nil {
		return "", err
	}

	return createdId.Hex(), nil
}
