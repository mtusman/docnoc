package pkg

import (
	"strings"
	"unicode/utf8"

	"github.com/fatih/color"
)

var (
	tO    = color.New(color.FgBlue).Add(color.Bold)
	cO    = color.New(color.FgGreen)
	iO    = color.New(color.FgYellow)
	width = 100
)

func printTitle(name string) {
	tO.Println(strings.ToUpper(name))
}

func printContainerName(name string, numErrs int) {
	keyMsg := "  \u2022 " + name
	space := strings.Repeat(".", Width-utf8.RuneCountInString(keyMsg))
	var emoji string
	if numErrs == 0 {
		emoji = "✅"
	} else {
		emoji = "😱"
	}
	cO.Println(keyMsg + space + emoji)
}

func printIssues(issues *Issues) {
	for _, issue := range *issues {
		printIssue(issue.message)
	}
}
func printIssue(message string) {
	iO.Println("\t", message)
}
