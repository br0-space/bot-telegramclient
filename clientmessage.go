package telegramclient

import (
	"regexp"
)

func Message(text string) MessageStruct {
	return MessageStruct{
		Text: text,
	}
}

func MarkdownMessage(text string) MessageStruct {
	return MessageStruct{
		Text:      text,
		ParseMode: "MarkdownV2",
	}
}

func Reply(text string, messageID int64) MessageStruct {
	return MessageStruct{
		Text:             text,
		ReplyToMessageID: messageID,
	}
}

func MarkdownReply(text string, messageID int64) MessageStruct {
	return MessageStruct{
		Text:             text,
		ReplyToMessageID: messageID,
		ParseMode:        "MarkdownV2",
	}
}

func Photo(photo string, caption string) MessageStruct {
	return MessageStruct{
		Photo:   photo,
		Caption: caption,
	}
}

func MarkdownPhoto(photo string, caption string) MessageStruct {
	return MessageStruct{
		Photo:     photo,
		Caption:   caption,
		ParseMode: "MarkdownV2",
	}
}

func EscapeMarkdown(text string) string {
	re := regexp.MustCompile("[_*\\[\\]()~`>#+\\-=|{}.!\\\\]")

	return re.ReplaceAllString(text, "\\$0")
}
