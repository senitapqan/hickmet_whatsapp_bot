package service

import "hickmet_whatapp_bot/models"

func (s *service) GetSessionByPhone(phone string) (models.Session, error) {
	return s.repos.GetSessionByPhone(phone)
}
