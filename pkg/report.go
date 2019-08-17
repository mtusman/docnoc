package pkg

import (
	"strings"
	"unicode/utf8"

	"github.com/fatih/color"
)

var (
	tO    = color.New(color.FgBlue).Add(color.Bold)
	cNO   = color.New(color.FgGreen)
	cIDO  = color.New(color.FgYellow)
	iO    = color.New(color.FgGreen)
	width = 100
)

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

func PrintContainerID(ID string) {
	cIDO.Println("    🐳 " + ID)
}

func PrintIssuesList(issues []*Issue) {
	for _, issue := range issues {
		PrintIssue(issue.message)
	}
}
func PrintIssue(message string) {
	iO.Println("\t😱", message)
}
