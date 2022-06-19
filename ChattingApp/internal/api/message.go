package api

type MessageRepository interface{}

type MessageService interface{}

type messageService struct {
	storage MessageRepository
}

func NewMessageService(messageRepo MessageRepository) MessageService {
	return &messageService{
		storage: messageRepo,
	}
}
