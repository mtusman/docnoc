package pkg

import "fmt"

type Issue struct {
	Message   string
	Processed bool
}

type Issues struct {
	containerID, containerName string
	IssuesList                 []*Issue
	ActionTaken                bool
}

func (i *Issues) MinMaxUtilisationIssue(cSV float64, cSN, cID string, underUtil bool) {
	var msg string
	issName := mapContainerStatNameToIssueName[cSN]
	if underUtil {
		msg = fmt.Sprintf("%s is under utilised at %0.2f", issName, cSV)
	} else {
		msg = fmt.Sprintf("%s is over utilised at %0.2f", issName, cSV)
	}

	for _, value := range ContainerStatNamePercs {
		if issName == value {
			msg += "%"
			(*i).IssuesList = append((*i).IssuesList, &Issue{Message: msg, Processed: false})
			return
		}
	}
	msg += "MB"
	(*i).IssuesList = append((*i).IssuesList, &Issue{Message: msg})
}
