package pkg

import (
	"fmt"

	"github.com/nlopes/slack"
)

func PostInitSlackMessage(webhook string) {
	msg := &slack.WebhookMessage{
		Username:  "docker",
		IconEmoji: ":whale:",
		Text:      "DocNoc has started scanning",
	}
	if err := slack.PostWebhook(webhook, msg); err != nil {
		fmt.Println("🔥: Can't post init message to slack", err)
	}
}
