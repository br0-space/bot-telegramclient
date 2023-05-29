package telegramclient

type ClientInterface interface {
	SendMessage(chatID int64, messageOut MessageStruct) error
}

// Create a struct that is accepted by Telegram's sendMessage endpoint
// https://core.telegram.org/bots/api#sendmessage

type MessageStruct struct {
	ChatID                int64  `json:"chat_id"`             //nolint:tagliatelle
	ReplyToMessageID      int64  `json:"reply_to_message_id"` //nolint:tagliatelle
	Text                  string `json:"text"`
	Photo                 string `json:"photo"`
	Caption               string `json:"caption"`
	ParseMode             string `json:"parse_mode"`               //nolint:tagliatelle
	DisableWebPagePreview bool   `json:"disable_web_page_preview"` //nolint:tagliatelle
	DisableNotification   bool   `json:"disable_notification"`     //nolint:tagliatelle
}

type MessageResponseBodyStruct struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int16  `json:"error_code"` //nolint:tagliatelle
	Description string `json:"description"`
}

type setWebhookURLResponse struct {
	Ok          bool   `json:"ok"`
	Result      bool   `json:"result"`
	ErrorCode   int    `json:"error_code"` //nolint:tagliatelle
	Description string `json:"description"`
}

type sendMessageResponse struct {
	Ok          bool   `json:"ok"`
	Result      bool   `json:"result"`
	ErrorCode   int    `json:"error_code"` //nolint:tagliatelle
	Description string `json:"description"`
}
