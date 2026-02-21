package service

import (
	"errors"
	"portfolio-send-message/model"
	"portfolio-send-message/storage"
)

type Service struct{
	storage storage.Storage
}

func (s *Service) AddMessage(messageInput model.MessageInput) error{

	if messageInput.Message == ""{
		return errors.New("no message passed u little hacker")
	}

	messageTitle := "Author: "
	if messageInput.Name == ""{
		messageTitle += "Anon"
	} else{
		messageTitle += messageInput.Name
	}

	var messageContent string;

	if messageInput.Email != ""{
		messageContent += "Email: "+messageInput.Email+"\n\n"
	}

	if messageInput.Discord != ""{
		messageContent += "Discord: "+messageInput.Discord+"\n\n"
	}

	messageContent += messageInput.Message

	err := s.storage.AddMessage(messageTitle, messageContent)
	return err
}

func NewService(storage storage.Storage) Service{
	return Service{storage: storage}
}
