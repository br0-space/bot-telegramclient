package telegramclient

func TestWebhookMessage(text string) WebhookMessageStruct {
	return WebhookMessageStruct{
		ID:      123, //nolint:gomnd
		From:    TestWebhookMessageUser(false),
		Chat:    TestWebhookMessageChat(),
		Text:    text,
		Date:    0,
		Photo:   []WebhookMessagePhotoStruct{},
		Caption: "",
	}
}

func TestWebhookMessageUser(isBot bool) WebhookMessageUserStruct {
	return WebhookMessageUserStruct{
		ID:           456, //nolint:gomnd
		IsBot:        isBot,
		FirstName:    "",
		LastName:     "",
		Username:     "Foobar",
		LanguageCode: "",
	}
}

func TestWebhookMessageChat() WebhookMessageChatStruct {
	return WebhookMessageChatStruct{
		ID:       789, //nolint:gomnd
		Type:     "",
		Username: "",
	}
}
