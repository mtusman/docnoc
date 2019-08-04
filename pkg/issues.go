package pkg

import "fmt"

type Issue struct {
	message string
}

type Issues []*Issue

func (i *Issues) AboveMaxUtilisationIssue(group string, amount float64) {
	message := fmt.Sprintf("%s is over utilised at %f", group, amount)
	*i = append(*i, &Issue{message})
}

func (i *Issues) AboveMinUtilisationIssue(group string, amount float64) {
	message := fmt.Sprintf("%s is under utilised at %f", group, amount)
	*i = append(*i, &Issue{message})
}
