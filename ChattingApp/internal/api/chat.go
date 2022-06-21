package api

import "chattingApp/internal/repository"

type ChatService interface {
	GetChats() (*[]Chat, error)
	CreateChat(chat *NewChatRequest) (string, error)
}

type chatService struct {
	storage repository.ChatRepository
}

func NewChatService(chatRepo repository.ChatRepository) ChatService {
	return &chatService{
		storage: chatRepo,
	}
}

func (s *chatService) GetChats() (*[]Chat, error) {
	chatEntities, err := s.storage.GetChats()
	if err != nil {
		return nil, err
	}

	chats := make([]Chat, len(*chatEntities))
	for i, chat := range *chatEntities {
		chats[i].Id = chat.Id.Hex()
		chats[i].Title = chat.Title
	}

	return &chats, nil
}

func (s *chatService) CreateChat(chat *NewChatRequest) (string, error) {
	chatEntity := &repository.Chat{
		Title: chat.Title,
	}

	createdId, err := s.storage.AddChat(chatEntity)
	if err != nil {
		return "", err
	}

	return createdId.Hex(), nil
}
