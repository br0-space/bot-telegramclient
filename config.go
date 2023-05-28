package telegramclient

type ConfigStruct struct {
	ApiKey              string
	WebhookURL          string
	BaseUrl             string
	EndpointSetWebhook  string
	EndpointSendMessage string
	EndpointSendPhoto   string
	ChatID              int64
}
