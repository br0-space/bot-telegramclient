package telegramclient

type ConfigStruct struct {
	APIKey              string
	WebhookURL          string
	BaseURL             string
	EndpointSetWebhook  string
	EndpointSendMessage string
	EndpointSendPhoto   string
	ChatID              int64
}
