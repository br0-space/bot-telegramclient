package telegramclient

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type WebhookHandlerInterface interface {
	ServeHTTP(res http.ResponseWriter, req *http.Request)
}

// Create a struct that mimics the webhook response body
// https://core.telegram.org/bots/api#update

type WebhookBodyStruct struct {
	Message WebhookMessageStruct `json:"message"`
}

type WebhookMessageStruct struct {
	ID      int64                       `json:"message_id"` //nolint:tagliatelle
	From    WebhookMessageUserStruct    `json:"from"`
	Chat    WebhookMessageChatStruct    `json:"chat"`
	Text    string                      `json:"text"`
	Date    int64                       `json:"date"`
	Photo   []WebhookMessagePhotoStruct `json:"photo"`
	Caption string                      `json:"caption"`
}

func (m WebhookMessageStruct) TextOrCaption() string {
	if len(m.Text) > 0 {
		return m.Text
	}

	if len(m.Caption) > 0 {
		return m.Caption
	}

	return ""
}

func (m WebhookMessageStruct) WordCount() int {
	// Match non-space character sequences.
	re := regexp.MustCompile(`\S+`)

	// Find all matches and return count.
	results := re.FindAllString(m.TextOrCaption(), -1)

	return len(results)
}

type WebhookMessageUserStruct struct {
	ID           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`     //nolint:tagliatelle
	FirstName    string `json:"first_name"` //nolint:tagliatelle
	LastName     string `json:"last_name"`  //nolint:tagliatelle
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"` //nolint:tagliatelle
}

func (u WebhookMessageUserStruct) UsernameOrName() string {
	if len(u.Username) > 0 {
		return "@" + u.Username
	}

	return strings.Trim(fmt.Sprintf("%s %s", u.FirstName, u.LastName), " ")
}

func (u WebhookMessageUserStruct) FirstnameOrUsername() string {
	if len(u.FirstName) > 0 {
		return u.FirstName
	}

	return "@" + u.Username
}

type WebhookMessageChatStruct struct {
	ID       int64  `json:"id"`
	Type     string `json:"type"`
	Username string `json:"username"`
}

type WebhookMessagePhotoStruct struct {
	FileID       string `json:"file_id"`        //nolint:tagliatelle
	FileUniqueID string `json:"file_unique_id"` //nolint:tagliatelle
	FileSize     int    `json:"file_size"`      //nolint:tagliatelle
	Width        int    `json:"width"`
	Height       int    `json:"height"`
}
