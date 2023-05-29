package telegramclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	logger "github.com/br0-space/bot-logger"
)

type Handler struct {
	log logger.Interface
	cfg *ConfigStruct
	fn  func(messageIn WebhookMessageStruct)
}

func NewHandler(
	config *ConfigStruct,
	fn func(messageIn WebhookMessageStruct),
) *Handler {
	handler := &Handler{
		log: logger.New(),
		cfg: config,
		fn:  fn,
	}
	handler.setWebhookURL()

	return handler
}

func (h *Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	h.log.Debugf("%s %s %s from %s", req.Method, req.URL, req.Proto, req.RemoteAddr)

	messageIn, status, err := h.parseRequest(req)
	if err != nil {
		h.log.Error(err)
		http.Error(res, err.Error(), status)

		return
	}

	h.fn(*messageIn)
}

func (h *Handler) parseRequest(req *http.Request) (*WebhookMessageStruct, int, error) {
	if req.Method != http.MethodPost {
		return nil, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed: %s (actual) != POST (expected)", req.Method)
	}

	body := &WebhookBodyStruct{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("unable to decode request body: %s", err.Error())
	}

	if h.cfg.ChatID != 0 && body.Message.Chat.ID != h.cfg.ChatID {
		return nil, http.StatusOK, fmt.Errorf("chat id mismatch: %d (actual) != %d (expected)", body.Message.Chat.ID, h.cfg.ChatID)
	}

	return &body.Message, 0, nil
}

func (h *Handler) setWebhookURL() {
	if h.cfg.WebhookURL == "" {
		h.log.Info("Not setting Telegram webhook URL")

		return
	}

	h.log.Info("Setting Telegram webhook URL to", h.cfg.WebhookURL)

	apiUrl := fmt.Sprintf(h.cfg.BaseUrl, h.cfg.ApiKey) + h.cfg.EndpointSetWebhook

	h.log.Debug("Sending POST request to", apiUrl)

	if resp, err := http.PostForm(apiUrl, url.Values{ //nolint:gosec
		"url": {h.cfg.WebhookURL},
	}); err != nil {
		h.log.Panic("Unable to set Telegram webhook URL:", err)
	} else {
		body := &setWebhookURLResponse{}
		if err = json.NewDecoder(resp.Body).Decode(body); err != nil {
			h.log.Fatal("Unable to decode response body:", err)
		}

		if !body.Ok {
			h.log.Fatal("Unable to set Telegram webhook URL:", body.Description)
		}

		h.log.Debug("Successfully set Telegram webhook URL")
	}
}
