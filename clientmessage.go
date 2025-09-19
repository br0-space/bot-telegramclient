package telegramclient

import (
	"regexp"
)

const parseModeMarkdown = "MarkdownV2"

// Message creates a basic MessageStruct with the provided text.
// It does not target any chat and has no special flags set.
func Message(text string) MessageStruct {
	return MessageStruct{
		ChatID:                0,
		ReplyToMessageID:      0,
		Text:                  text,
		Photo:                 "",
		Caption:               "",
		ParseMode:             "",
		DisableWebPagePreview: false,
		DisableNotification:   false,
	}
}

// MessageToChat creates a basic text message targeting a specific chat.
func MessageToChat(text string, chatID int64) MessageStruct {
	m := Message(text)
	m.ChatID = chatID
	return m
}

// MarkdownMessage creates a text message using the MarkdownV2 parse mode.
// Use EscapeMarkdown to escape user-provided text when needed.
func MarkdownMessage(text string) MessageStruct {
	message := Message(text)
	message.ParseMode = parseModeMarkdown

	return message
}

// MarkdownMessageToChat creates a MarkdownV2 message for a specific chat.
func MarkdownMessageToChat(text string, chatID int64) MessageStruct {
	m := MarkdownMessage(text)
	m.ChatID = chatID
	return m
}

// Reply creates a message that replies to a specific message ID.
// It does not set a chat target by itself.
func Reply(text string, messageID int64) MessageStruct {
	message := Message(text)
	message.ReplyToMessageID = messageID

	return message
}

// ReplyToChat creates a reply to a specific message in a specific chat.
func ReplyToChat(text string, messageID int64, chatID int64) MessageStruct {
	m := Reply(text, messageID)
	m.ChatID = chatID
	return m
}

// MarkdownReply creates a reply message using the MarkdownV2 parse mode.
// Use EscapeMarkdown to escape user-provided text when needed.
func MarkdownReply(text string, messageID int64) MessageStruct {
	message := Reply(text, messageID)
	message.ParseMode = parseModeMarkdown

	return message
}

// MarkdownReplyToChat creates a MarkdownV2 reply in a specific chat.
func MarkdownReplyToChat(text string, messageID int64, chatID int64) MessageStruct {
	m := MarkdownReply(text, messageID)
	m.ChatID = chatID
	return m
}

// Photo creates a message that sends a photo with an optional caption.
// The caption is also assigned to the Text field for compatibility.
func Photo(photo string, caption string) MessageStruct {
	message := Message(caption)
	message.Photo = photo
	message.Caption = caption

	return message
}

// PhotoToChat creates a photo message with caption for a specific chat.
func PhotoToChat(photo string, caption string, chatID int64) MessageStruct {
	m := Photo(photo, caption)
	m.ChatID = chatID
	return m
}

// MarkdownPhoto creates a photo message with caption using the MarkdownV2 parse mode.
// Use EscapeMarkdown to escape user-provided caption when needed.
func MarkdownPhoto(photo string, caption string) MessageStruct {
	message := Photo(photo, caption)
	message.ParseMode = parseModeMarkdown

	return message
}

// MarkdownPhotoToChat creates a MarkdownV2 photo message for a specific chat.
func MarkdownPhotoToChat(photo string, caption string, chatID int64) MessageStruct {
	m := MarkdownPhoto(photo, caption)
	m.ChatID = chatID
	return m
}

// EscapeMarkdown escapes characters that have special meaning in Telegram MarkdownV2.
// It helps prevent unintended formatting when sending user-generated content.
func EscapeMarkdown(text string) string {
	re := regexp.MustCompile("[_*\\[\\]()~`>#+\\-=|{}.!\\\\]")

	return re.ReplaceAllString(text, "\\$0")
}
