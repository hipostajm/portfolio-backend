package storage

type Storage interface{
	AddMessage(messageTitle string, messageContent string) error
}
