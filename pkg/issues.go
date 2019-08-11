package pkg

import "fmt"

type Issue struct {
	message string
}

type Issues map[string][]*Issue

func (i *Issues) AboveMaxUtilisationIssue(group string, amount float64, containerID string) {
	message := fmt.Sprintf("%s is over utilised at %0.2f%%", group, amount)
	(*i)[containerID] = append((*i)[containerID], &Issue{message})
}

func (i *Issues) AboveMinUtilisationIssue(group string, amount float64, containerID string) {
	message := fmt.Sprintf("%s is under utilised at %0.2f%%", group, amount)
	(*i)[containerID] = append((*i)[containerID], &Issue{message})
}
