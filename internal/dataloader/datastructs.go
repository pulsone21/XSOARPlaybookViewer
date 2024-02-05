package dataloader

import (
	"fmt"
	"strings"
)

type Playbook struct {
	Id        string
	Name      string
	FileName  string
	Inputs    []PlaybookInput
	Outputs   []PlaybookOutput
	NextTasks []PlaybookTask
}

type PlaybookInput struct {
	Name        string
	Required    bool
	Description string
}

type PlaybookOutput struct {
	Name        string
	Description string
	Type        string
}

type PlaybookTask struct {
	Id        string
	NextTasks map[string][]string
	Type      PlaybookTaskType
	Name      string
	Task      map[string]interface{}
}

type PlaybookTaskType string

const (
	playbook   PlaybookTaskType = "playbook"
	Start      PlaybookTaskType = "start"
	Automation PlaybookTaskType = "automation"
	Condition  PlaybookTaskType = "condition"
	Title      PlaybookTaskType = "title"
)

func (t *PlaybookTask) GetFileName() (string, error) {
	if t.Type == playbook || t.Type == Automation {
		//only then we have a file associated with it
		fName := strings.ReplaceAll(t.Name, " ", "_")
		return fName, nil
	}
	return "", fmt.Errorf("no file associated with taskType %s", t.Type)
}

func toTaskType(str string) PlaybookTaskType {
	switch str {
	case "playbook":
		return playbook
	case "start":
		return Start
	case "automation":
		return Automation
	case "condition":
		return Condition
	case "title":
		return Title
	}
	return "unknown"
}
