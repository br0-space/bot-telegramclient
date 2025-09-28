# bot-telegramclient

A small Go library for sending messages to Telegram and handling Telegram Bot webhooks. It provides:

- A simple client to send text and photo messages
- Helpers to build messages (including MarkdownV2-safe helpers)
- An HTTP handler to process incoming webhook updates from Telegram

This repository is intended to be used as a building block inside Telegram bots or services that need to interact with Telegram.

## Status

[![Build](https://github.com/br0-space/bot-telegramclient/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/br0-space/bot-telegramclient/actions/workflows/build.yml)
[![Test](https://github.com/br0-space/bot-telegramclient/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/br0-space/bot-telegramclient/actions/workflows/test.yml)
[![Lint](https://github.com/br0-space/bot-telegramclient/actions/workflows/lint.yml/badge.svg?branch=main)](https://github.com/br0-space/bot-telegramclient/actions/workflows/lint.yml)
[![Staticcheck](https://github.com/br0-space/bot-telegramclient/actions/workflows/staticcheck.yml/badge.svg?branch=main)](https://github.com/br0-space/bot-telegramclient/actions/workflows/staticcheck.yml)
[![Vet](https://github.com/br0-space/bot-telegramclient/actions/workflows/vet.yml/badge.svg?branch=main)](https://github.com/br0-space/bot-telegramclient/actions/workflows/vet.yml)
[![CodeQL](https://github.com/br0-space/bot-telegramclient/actions/workflows/codeql-analysis.yml/badge.svg?branch=main)](https://github.com/br0-space/bot-telegramclient/actions/workflows/codeql-analysis.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/br0-space/bot-telegramclient.svg)](https://pkg.go.dev/github.com/br0-space/bot-telegramclient)
[![Go Report Card](https://goreportcard.com/badge/github.com/br0-space/bot-telegramclient)](https://goreportcard.com/report/github.com/br0-space/bot-telegramclient)


## Installation

Go 1.20+ is recommended.

```
go get github.com/br0-space/bot-telegramclient
```


## Quick start

Send a simple message to a chat:

```go
package main

import (
    "log"

    telegramclient "github.com/br0-space/bot-telegramclient"
)

func main() {
    cfg := telegramclient.ConfigStruct{
        APIKey:              "<BOT_TOKEN>",
        BaseURL:             "https://api.telegram.org/bot%s/",
        EndpointSendMessage: "sendMessage",
        EndpointSendPhoto:   "sendPhoto",
        // Optional: lock webhooks to a chat or configure webhook URL later
        ChatID: 0,
    }

    c := telegramclient.NewClient(cfg)

    if err := c.SendMessage(1234567890, telegramclient.Message("Hello, Telegram!")); err != nil {
        log.Fatal(err)
    }
}
```

Replace `1234567890` with your destination chat ID.


## Working with MarkdownV2

Telegram uses a strict MarkdownV2 dialect. This library offers helpers to set the parse mode and escape user-generated input safely.

```go
msg := telegramclient.MarkdownMessage("*Bold* _Italic_ \nInline `code`")
// If you interpolate user text, escape it to avoid formatting issues:
user := "Price is 10% (great!)"
escaped := telegramclient.EscapeMarkdown(user)
msg.Text = "User said: " + escaped

// Send to a chat
_ = c.SendMessage(1234567890, msg)
```

Photo with MarkdownV2 caption:

```go
photo := telegramclient.MarkdownPhoto("https://example.com/cat.jpg", "Look at this _cat_")
_ = c.SendMessage(1234567890, photo)
```

Replying to a message in a chat:

```go
reply := telegramclient.ReplyToChat("Thanks!", 42, 1234567890)
_ = c.SendMessage(1234567890, reply)
```


## Handling Telegram webhooks

Use the provided HTTP handler to accept Telegram updates. The handler validates the method, optionally checks an expected chat ID, and invokes your callback with a parsed message structure.

```go
package main

import (
    "log"
    "net/http"

    telegramclient "github.com/br0-space/bot-telegramclient"
)

func main() {
    cfg := telegramclient.ConfigStruct{
        APIKey:             "<BOT_TOKEN>",
        BaseURL:            "https://api.telegram.org/bot%s/",
        EndpointSetWebhook: "setWebhook",
        // Optional: if set, only messages from this chat ID are accepted
        ChatID: 1234567890,
        // Optional: if set, the handler will set the Telegram webhook to this URL on startup
        WebhookURL: "https://your.domain/telegram/webhook",
    }

    handler := telegramclient.NewHandler(&cfg, func(in telegramclient.WebhookMessageStruct) {
        log.Printf("Got message from @%s: %s", in.From.Username, in.Text)
        // Do something with the incoming message...
    })

    http.Handle("/telegram/webhook", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Notes:
- If `WebhookURL` is set, the handler will call Telegram's `setWebhook` endpoint during initialization.
- If `ChatID` is non-zero, messages from other chats are ignored with HTTP 200 to keep Telegram happy.


## Configuration reference

These fields are used across the client and handler:

- APIKey: Telegram Bot token (without the `bot` prefix; the BaseURL format will inject it)
- WebhookURL: Public HTTPS endpoint Telegram should call; if set, `NewHandler` will set it
- BaseURL: Base API URL, e.g. `https://api.telegram.org/bot%s/` (the `%s` is replaced with APIKey)
- EndpointSetWebhook: Path for setting webhooks, usually `setWebhook`
- EndpointSendMessage: Path for sending messages, usually `sendMessage`
- EndpointSendPhoto: Path for sending photos, usually `sendPhoto`
- ChatID: Optional chat ID filter for incoming webhooks


## Message helpers

Factory functions that produce `MessageStruct`:

- Message(text)
- MessageToChat(text, chatID)
- MarkdownMessage(text)
- MarkdownMessageToChat(text, chatID)
- Reply(text, messageID)
- ReplyToChat(text, messageID, chatID)
- MarkdownReply(text, messageID)
- MarkdownReplyToChat(text, messageID, chatID)
- Photo(photoURLOrFileID, caption)
- PhotoToChat(photoURLOrFileID, caption, chatID)
- MarkdownPhoto(photoURLOrFileID, caption)
- MarkdownPhotoToChat(photoURLOrFileID, caption, chatID)
- EscapeMarkdown(text) — escapes Telegram MarkdownV2 special characters

Send with:

```go
client.SendMessage(chatID, msg)
```

The client automatically selects `sendPhoto` when `Photo` is non-empty; otherwise it uses `sendMessage`.


## Testing and mocks

For tests, you can implement `ClientInterface` or use your own stub. This repo includes a `mockclient.go` that can be adapted for testing scenarios.

There are also helpers for constructing example webhook messages in `webhookmessage.go` for tests.


## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
