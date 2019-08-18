package pkg

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/nlopes/slack"

	"github.com/fatih/color"
)

var (
	tO    = color.New(color.FgBlue).Add(color.Bold)
	cNO   = color.New(color.FgGreen)
	cIDO  = color.New(color.FgYellow)
	IO    = color.New(color.FgGreen)
	width = 100
)

func ProcessReportForApp(key string, numErrs int) {
	PrintContainerName(key, numErrs)

}

func PrintTitle(name string) {
	tO.Println(strings.ToUpper(name))
}

func PrintContainerName(name string, numErrs int) {
	keyMsg := "  \u2022 " + name
	space := strings.Repeat(".", width-utf8.RuneCountInString(keyMsg))
	var emoji string
	if numErrs == 0 {
		emoji = "✅"
	} else {
		emoji = "😱"
	}
	cNO.Println(keyMsg + space + emoji)
}

func PrintIssuesList(dN *DocNoc, cN, cID string, issues []*Issue) {
	cIDO.Println("    🐳 " + cID)
	for _, issue := range issues {
		PrintIssue(issue.Message)
		slackWebhook := dN.DocNocConfig.SlackWebhook
		if slackWebhook != "" && !issue.Processed {
			slack.PostWebhook(slackWebhook, &slack.WebhookMessage{
				Username:  Username,
				IconEmoji: IconEmoji,
				Attachments: []slack.Attachment{
					slack.Attachment{
						Title:      fmt.Sprintf(":package: Container %s", cN),
						Text:       fmt.Sprintf("Container `%s` experienced the following issue: %s", cN, issue.Message),
						Footer:     cID,
						Color:      "danger",
						MarkdownIn: []string{"text"},
					},
				},
			})
		}
		issue.Processed = true
	}
}

func PrintIssue(message string) {
	IO.Println("\t😱", message)
}
