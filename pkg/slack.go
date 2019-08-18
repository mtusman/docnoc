package pkg

import (
	"fmt"

	"github.com/nlopes/slack"
)

var (
	Username  = "docker"
	IconEmoji = ":whale:"
)

func PostInitSlackMessage(webhook string) {
	msg := &slack.WebhookMessage{
		Username:  Username,
		IconEmoji: IconEmoji,
		Text:      "DocNoc has started scanning",
	}
	if err := slack.PostWebhook(webhook, msg); err != nil {
		fmt.Println("🔥: Can't post init message to slack", err)
	}
}
