package telegramclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	logger "github.com/br0-space/bot-logger"
	"net/http"
)

type Client struct {
	Log logger.Interface
	Cfg ConfigStruct
}

func NewClient(config ConfigStruct) *Client {
	return &Client{
		Log: logger.New(),
		Cfg: config,
	}
}

func (c Client) SendMessage(chatID int64, message MessageStruct) error {
	switch {
	case message.Photo != "":
		c.Log.Debugf("Sending photo: %s", message.Photo)
	default:
		c.Log.Debugf("Sending message: %s", message.Text)
	}

	message.ChatID = chatID

	url := c.url(message)
	requestBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	c.Log.Debugf("Sending POST request to %s", url)

	response, err := http.Post(url, "application/json", bytes.NewBuffer(requestBytes))
	if err != nil {
		return err
	}

	if response.StatusCode == http.StatusOK {
		c.Log.Debug("Successfully sent message to Telegram")

		return nil
	}

	responseBody := &sendMessageResponse{}
	if err = json.NewDecoder(response.Body).Decode(responseBody); err != nil {
		return fmt.Errorf("SendMessage failed with %s: unable to decode response body", response.Status)
	}
	return fmt.Errorf("SendMessage failed with %d: %s", responseBody.ErrorCode, responseBody.Description)
}

func (c Client) url(message MessageStruct) string {
	switch {
	case message.Photo != "":
		return fmt.Sprintf(c.Cfg.BaseUrl, c.Cfg.ApiKey) + c.Cfg.EndpointSendPhoto
	default:
		return fmt.Sprintf(c.Cfg.BaseUrl, c.Cfg.ApiKey) + c.Cfg.EndpointSendMessage
	}
}
