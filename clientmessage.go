package telegramclient

import (
	"regexp"
)

const parseModeMarkdown = "MarkdownV2"

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

func MarkdownMessage(text string) MessageStruct {
	message := Message(text)
	message.ParseMode = parseModeMarkdown

	return message
}

func Reply(text string, messageID int64) MessageStruct {
	message := Message(text)
	message.ReplyToMessageID = messageID

	return message
}

func MarkdownReply(text string, messageID int64) MessageStruct {
	message := Reply(text, messageID)
	message.ParseMode = parseModeMarkdown

	return message
}

func Photo(photo string, caption string) MessageStruct {
	message := Message(caption)
	message.Photo = photo
	message.Caption = caption

	return message
}

func MarkdownPhoto(photo string, caption string) MessageStruct {
	message := Photo(photo, caption)
	message.ParseMode = parseModeMarkdown

	return message
}

func EscapeMarkdown(text string) string {
	re := regexp.MustCompile("[_*\\[\\]()~`>#+\\-=|{}.!\\\\]")

	return re.ReplaceAllString(text, "\\$0")
}
