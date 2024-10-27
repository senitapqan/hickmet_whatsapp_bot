package api_client

import (
	"github.com/rs/zerolog/log"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func (t *twilioClient) SendMessage(message string, phone string) error {
	params := &openapi.CreateMessageParams{}
	params.SetTo("whatsapp:" + phone)
	log.Info().Msg("whatsapp:" + phone)     // Номер получателя
	params.SetFrom("whatsapp:+14155238886") // Твой номер в Twilio
	params.SetBody(message)                 // Сообщение

	_, err := t.client.Api.CreateMessage(params)
	if err != nil {
		log.Info().Msg(err.Error())
	}
	return err
}
