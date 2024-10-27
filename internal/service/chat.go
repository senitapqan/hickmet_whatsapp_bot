package service

import (
	"hickmet_whatapp_bot/consts"

	"github.com/rs/zerolog/log"
)

func (s *service) Response(phone string, content string) error {
	session, err := s.repos.GetSessionByPhone(phone)
	if err != nil {
		log.Info().Msg("error was detected: " + err.Error())
		return err
	}

	log.Info().Msg("session status: " + session.Status)
	if session.Status == consts.Start {
		err = s.twilio.SendMessage(consts.Start_message, phone)
		if err != nil {
			return err
		}
	}

	return nil
}
