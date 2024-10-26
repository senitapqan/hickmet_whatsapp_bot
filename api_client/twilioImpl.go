package api_client

import openapi "github.com/twilio/twilio-go/rest/api/v2010"

func (t *twilioClient) SendMessage(message string) error {
	params := &openapi.CreateMessageParams{}
	params.SetTo("87713428110")   // Номер получателя
	params.SetFrom("87086815386") // Твой номер в Twilio
	params.SetBody("body")        // Сообщение

	_, err := t.client.Api.CreateMessage(params)
	return err
}
