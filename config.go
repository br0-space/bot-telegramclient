package telegramclient

type ConfigStruct struct {
	APIKey              string
	WebhookURL          string
	BaseURL             string
	EndpointSetWebhook  string
	EndpointSendMessage string
	EndpointSendPhoto   string
	// ChatIDs allows filtering incoming webhook messages to specific chats.
	// If non-empty, only messages from these chat IDs are processed.
	ChatIDs             []int64
	// ChatID is kept for backward compatibility; used only if ChatIDs is empty.
	ChatID              int64
}
