package pkg

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/nlopes/slack"

	"github.com/fatih/color"
)

var (
	// tO is used to print title output
	tO = color.New(color.FgBlue).Add(color.Bold)
	// cNO is used to print container name output
	cNO = color.New(color.FgGreen)
	// cIDO is used to print container ID output
	cIDO = color.New(color.FgYellow)
	// IO is used to print general text
	IO = color.New(color.FgGreen)
	// width used to calculate to end line formatting
	width = 100
)

// ProcessReportForApp is used to call PrintContainerName
func ProcessReportForApp(key string, numErrs int) {
	PrintContainerName(key, numErrs)

}

// PrintTitle is used to print the title
func PrintTitle(name string) {
	tO.Println(strings.ToUpper(name))
}

// PrintContainerName is used to print the container name with the appropriate emoji :)
// and format it correctly
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

// PrintIssuesList is used to print all the issues associated with a particular container
// and send a slack message.
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
		// We don't want to send the same slack message more than once!
		issue.Processed = true
	}
}

// PrintIssue is used to print a single issue
func PrintIssue(message string) {
	IO.Println("\t😱", message)
}
