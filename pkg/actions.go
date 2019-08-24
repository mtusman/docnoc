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
		fmt.Println("🔥: Can't post init message to slack. Operating in headless state", err)
	}
}

func PostActionMessage(webhook, cN, cID, action string, errType bool) {
	if errType {
		IO.Println(fmt.Sprintf("\t🔥 Failed to %s container with ID: %s", action, cID))
	} else {
		IO.Println(fmt.Sprintf("\t🚒 %s container with ID: %s", action, cID))
	}

	if webhook != "" {
		var text, color string
		if errType {
			text = fmt.Sprintf("Failed to %s container `%s`", action, cN)
			color = "warning"
		} else {
			text = fmt.Sprintf("%s container `%s`", action, cN)
			color = "good"
		}
		slack.PostWebhook(webhook, &slack.WebhookMessage{
			Username:  Username,
			IconEmoji: IconEmoji,
			Attachments: []slack.Attachment{
				slack.Attachment{
					Title:      fmt.Sprintf(":package: Container %s", cN),
					Text:       text,
					Footer:     cID,
					Color:      color,
					MarkdownIn: []string{"text"},
				},
			},
		})
	}

}
