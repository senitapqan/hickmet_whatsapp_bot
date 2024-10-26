package service

import (
	"hickmet_whatapp_bot/internal/repository"
	"hickmet_whatapp_bot/models"
)

type Service interface {
	GetUserByIIN(iin string) (models.User, error)
}

type service struct {
	repos repository.Repository
}

func NewService(repos repository.Repository) Service {
	return &service{
		repos: repos,
	}
}
