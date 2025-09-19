package telegramclient_test

import (
	"testing"

	telegramclient "github.com/br0-space/bot-telegramclient"
)

func TestMessage(t *testing.T) {
	t.Parallel()

	text := "hello"
	m := telegramclient.Message(text)

	if m.Text != text {
		t.Errorf("Message.Text = %q, want %q", m.Text, text)
	}

	if m.ChatID != 0 {
		t.Errorf("Message.ChatID = %d, want 0", m.ChatID)
	}

	if m.ReplyToMessageID != 0 {
		t.Errorf("Message.ReplyToMessageID = %d, want 0", m.ReplyToMessageID)
	}

	if m.Photo != "" {
		t.Errorf("Message.Photo = %q, want empty", m.Photo)
	}

	if m.Caption != "" {
		t.Errorf("Message.Caption = %q, want empty", m.Caption)
	}

	if m.ParseMode != "" {
		t.Errorf("Message.ParseMode = %q, want empty", m.ParseMode)
	}

	if m.DisableWebPagePreview {
		t.Errorf("Message.DisableWebPagePreview = %v, want false", m.DisableWebPagePreview)
	}

	if m.DisableNotification {
		t.Errorf("Message.DisableNotification = %v, want false", m.DisableNotification)
	}
}

func TestMessageToChat(t *testing.T) {
	t.Parallel()

	text := "hello"
	chatID := int64(111)

	m := telegramclient.MessageToChat(text, chatID)
	if m.Text != text {
		t.Errorf("MessageToChat.Text = %q, want %q", m.Text, text)
	}

	if m.ChatID != chatID {
		t.Errorf("MessageToChat.ChatID = %d, want %d", m.ChatID, chatID)
	}
}

func TestMarkdownMessage(t *testing.T) {
	t.Parallel()

	text := "hello"
	m := telegramclient.MarkdownMessage(text)

	if m.Text != text {
		t.Errorf("MarkdownMessage.Text = %q, want %q", m.Text, text)
	}

	if m.ParseMode != "MarkdownV2" {
		t.Errorf("MarkdownMessage.ParseMode = %q, want MarkdownV2", m.ParseMode)
	}
}

func TestMarkdownMessageToChat(t *testing.T) {
	t.Parallel()

	text := "hello"
	chatID := int64(222)

	m := telegramclient.MarkdownMessageToChat(text, chatID)
	if m.Text != text {
		t.Errorf("MarkdownMessageToChat.Text = %q, want %q", m.Text, text)
	}

	if m.ChatID != chatID {
		t.Errorf("MarkdownMessageToChat.ChatID = %d, want %d", m.ChatID, chatID)
	}

	if m.ParseMode != "MarkdownV2" {
		t.Errorf("MarkdownMessageToChat.ParseMode = %q, want MarkdownV2", m.ParseMode)
	}
}

func TestReply(t *testing.T) {
	t.Parallel()

	text := "reply"
	msgID := int64(12345)
	m := telegramclient.Reply(text, msgID)

	if m.Text != text {
		t.Errorf("Reply.Text = %q, want %q", m.Text, text)
	}

	if m.ReplyToMessageID != msgID {
		t.Errorf("Reply.ReplyToMessageID = %d, want %d", m.ReplyToMessageID, msgID)
	}

	if m.ParseMode != "" {
		t.Errorf("Reply.ParseMode = %q, want empty", m.ParseMode)
	}
}

func TestReplyToChat(t *testing.T) {
	t.Parallel()

	text := "reply"
	chatID := int64(333)
	msgID := int64(444)

	m := telegramclient.ReplyToChat(text, msgID, chatID)
	if m.Text != text {
		t.Errorf("ReplyToChat.Text = %q, want %q", m.Text, text)
	}

	if m.ChatID != chatID {
		t.Errorf("ReplyToChat.ChatID = %d, want %d", m.ChatID, chatID)
	}

	if m.ReplyToMessageID != msgID {
		t.Errorf("ReplyToChat.ReplyToMessageID = %d, want %d", m.ReplyToMessageID, msgID)
	}
}

func TestMarkdownReply(t *testing.T) {
	t.Parallel()

	text := "reply"
	msgID := int64(54321)
	m := telegramclient.MarkdownReply(text, msgID)

	if m.Text != text {
		t.Errorf("MarkdownReply.Text = %q, want %q", m.Text, text)
	}

	if m.ReplyToMessageID != msgID {
		t.Errorf("MarkdownReply.ReplyToMessageID = %d, want %d", m.ReplyToMessageID, msgID)
	}

	if m.ParseMode != "MarkdownV2" {
		t.Errorf("MarkdownReply.ParseMode = %q, want MarkdownV2", m.ParseMode)
	}
}

func TestMarkdownReplyToChat(t *testing.T) {
	t.Parallel()

	text := "reply"
	chatID := int64(555)
	msgID := int64(666)

	m := telegramclient.MarkdownReplyToChat(text, msgID, chatID)
	if m.Text != text {
		t.Errorf("MarkdownReplyToChat.Text = %q, want %q", m.Text, text)
	}

	if m.ChatID != chatID {
		t.Errorf("MarkdownReplyToChat.ChatID = %d, want %d", m.ChatID, chatID)
	}

	if m.ReplyToMessageID != msgID {
		t.Errorf("MarkdownReplyToChat.ReplyToMessageID = %d, want %d", m.ReplyToMessageID, msgID)
	}

	if m.ParseMode != "MarkdownV2" {
		t.Errorf("MarkdownReplyToChat.ParseMode = %q, want MarkdownV2", m.ParseMode)
	}
}

func TestPhoto(t *testing.T) {
	t.Parallel()

	photo := "file_id_abc"
	caption := "look at this"
	m := telegramclient.Photo(photo, caption)

	if m.Photo != photo {
		t.Errorf("Photo.Photo = %q, want %q", m.Photo, photo)
	}

	if m.Caption != caption {
		t.Errorf("Photo.Caption = %q, want %q", m.Caption, caption)
	}

	if m.Text != caption {
		t.Errorf("Photo.Text = %q, want %q", m.Text, caption)
	}

	if m.ParseMode != "" {
		t.Errorf("Photo.ParseMode = %q, want empty", m.ParseMode)
	}
}

func TestPhotoToChat(t *testing.T) {
	t.Parallel()

	photo := "file_id_photo"
	caption := "a caption"
	chatID := int64(777)

	m := telegramclient.PhotoToChat(photo, caption, chatID)
	if m.Photo != photo {
		t.Errorf("PhotoToChat.Photo = %q, want %q", m.Photo, photo)
	}

	if m.Caption != caption {
		t.Errorf("PhotoToChat.Caption = %q, want %q", m.Caption, caption)
	}

	if m.Text != caption {
		t.Errorf("PhotoToChat.Text = %q, want %q", m.Text, caption)
	}

	if m.ChatID != chatID {
		t.Errorf("PhotoToChat.ChatID = %d, want %d", m.ChatID, chatID)
	}
}

func TestMarkdownPhoto(t *testing.T) {
	t.Parallel()

	photo := "file_id_xyz"
	caption := "caption here"
	m := telegramclient.MarkdownPhoto(photo, caption)

	if m.Photo != photo {
		t.Errorf("MarkdownPhoto.Photo = %q, want %q", m.Photo, photo)
	}

	if m.Caption != caption {
		t.Errorf("MarkdownPhoto.Caption = %q, want %q", m.Caption, caption)
	}

	if m.Text != caption {
		t.Errorf("MarkdownPhoto.Text = %q, want %q", m.Text, caption)
	}

	if m.ParseMode != "MarkdownV2" {
		t.Errorf("MarkdownPhoto.ParseMode = %q, want MarkdownV2", m.ParseMode)
	}
}

func TestMarkdownPhotoToChat(t *testing.T) {
	t.Parallel()

	photo := "file_id_markdown"
	caption := "md caption"
	chatID := int64(888)

	m := telegramclient.MarkdownPhotoToChat(photo, caption, chatID)
	if m.Photo != photo {
		t.Errorf("MarkdownPhotoToChat.Photo = %q, want %q", m.Photo, photo)
	}

	if m.Caption != caption {
		t.Errorf("MarkdownPhotoToChat.Caption = %q, want %q", m.Caption, caption)
	}

	if m.Text != caption {
		t.Errorf("MarkdownPhotoToChat.Text = %q, want %q", m.Text, caption)
	}

	if m.ChatID != chatID {
		t.Errorf("MarkdownPhotoToChat.ChatID = %d, want %d", m.ChatID, chatID)
	}

	if m.ParseMode != "MarkdownV2" {
		t.Errorf("MarkdownPhotoToChat.ParseMode = %q, want MarkdownV2", m.ParseMode)
	}
}

// EscapeMarkdown tests and helpers placed last to mirror function order

type escapeMarkdownTest struct {
	in  string
	out string
}

var escapeMarkdownTests = []escapeMarkdownTest{
	{"_", "\\_"},
	{"*", "\\*"},
	{"[", "\\["},
	{"]", "\\]"},
	{"(", "\\("},
	{")", "\\)"},
	{"~", "\\~"},
	{"`", "\\`"},
	{"<", "<"}, // not escaped by our regex
	{">", "\\>"},
	{"#", "\\#"},
	{"+", "\\+"},
	{"-", "\\-"},
	{"=", "\\="},
	{"|", "\\|"},
	{"{", "\\{"},
	{"}", "\\}"},
	{".", "\\."},
	{"!", "\\!"},
	{"\\", "\\\\"},
}

func TestEscapeMarkdown(t *testing.T) {
	t.Parallel()

	for _, test := range escapeMarkdownTests {
		if out := telegramclient.EscapeMarkdown(test.in); out != test.out {
			t.Errorf("EscapeMarkdown(%q) = %q, want %q", test.in, out, test.out)
		}
	}
}
