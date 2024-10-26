package api_client

import (
	"os"

	"github.com/twilio/twilio-go"
)

type TwilioClient interface {
	SendMessage(message string) error
}

type twilioClient struct {
	client *twilio.RestClient
}

func createTwilioClient() TwilioClient {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_ACCOUNT_SID"),
		Password: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	return &twilioClient{client: client}
}
