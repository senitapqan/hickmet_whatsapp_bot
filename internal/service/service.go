package service

import (
	"hickmet_whatapp_bot/api_client"
	"hickmet_whatapp_bot/internal/repository"
	"hickmet_whatapp_bot/models"
)

type Service interface {
	GetUserByIIN(iin string) (models.User, error)

	GetSessionByPhone(phone string) (models.Session, error)

	Response(phone string, content string) error
}

type service struct {
	repos  repository.Repository
	twilio api_client.TwilioClient
}

func NewService(repos repository.Repository, twilio api_client.TwilioClient) Service {
	return &service{
		repos:  repos,
		twilio: twilio,
	}
}
