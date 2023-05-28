package telegramclient

func TestWebhookMessage(text string) WebhookMessageStruct {
	return WebhookMessageStruct{
		ID:   123,
		From: TestWebhookMessageUser(false),
		Chat: TestWebhookMessageChat(),
		Text: text,
	}
}

func TestWebhookMessageUser(isBot bool) WebhookMessageUserStruct {
	return WebhookMessageUserStruct{
		ID:       456,
		IsBot:    isBot,
		Username: "Foobar",
	}
}

func TestWebhookMessageChat() WebhookMessageChatStruct {
	return WebhookMessageChatStruct{
		ID: 789,
	}
}
