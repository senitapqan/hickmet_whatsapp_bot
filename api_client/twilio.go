package api_client

import (
	"os"

	"github.com/twilio/twilio-go"
)

type TwilioClient interface {
	SendMessage(message string, phone string) error
}

type twilioClient struct {
	client *twilio.RestClient
}

func CreateTwilioClient() TwilioClient {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_ACCOUNT_SID"),
		Password: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	return &twilioClient{client: client}
}
