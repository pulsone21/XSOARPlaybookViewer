package dataloader

import (
	"fmt"
	"strings"
)

type Playbook struct {
	Id       string           `json:"id"`
	Name     string           `json:"name"`
	FileName string           `json:"file_name"`
	Inputs   []PlaybookInput  `json:"inputs"`
	Outputs  []PlaybookOutput `json:"outputs"`
	Tasks    []Task           `json:"tasks"`
}

type PlaybookInput struct {
	Name        string `json:"name"`
	Required    bool   `json:"required"`
	Description string `json:"description"`
}

type PlaybookOutput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

type Task interface {
	GetNextTask() map[string][]string
	GetFileName() string
	CreateEdges() *[]Edge
	CreateNode() *Node
}

type StartTask struct {
	Id        string              `json:"id"`
	NextTasks map[string][]string `json:"next_tasks"`
	Type      PlaybookTaskType    `json:"type"`
	TaskId    string              `json:"task_id"`
}

func CreateStartTask(id, taskId string, nT map[string][]string) *StartTask {
	return &StartTask{
		Id:        id,
		TaskId:    taskId,
		Type:      Start,
		NextTasks: nT,
	}
}

func (t StartTask) GetNextTask() map[string][]string {
	return t.NextTasks
}

func (t StartTask) GetFileName() string {
	// Root task didn't has a file assossiated
	return ""
}

func (t StartTask) CreateNode() *Node {
	return CreateNode(t.TaskId, t.Type, nil)
}

func (t StartTask) CreateEdges() *[]Edge {
	edges := []Edge{}
	for k, vals := range t.NextTasks {
		for _, val := range vals {
			e := Edge{
				Id:     fmt.Sprintf("%s-%s", t.Id, val),
				Source: t.Id,
				Target: val,
				Label:  k,
			}
			edges = append(edges, e)
		}
	}
	return &edges
}

type TitleTask struct {
	Id        string              `json:"id"`
	NextTasks map[string][]string `json:"next_tasks"`
	Type      PlaybookTaskType    `json:"type"`
	Name      string              `json:"name"`
	TaskId    string              `json:"task_id"`
}

func (t TitleTask) GetNextTask() map[string][]string {
	return t.NextTasks
}

func CreateTitleTask(id, taskId string, nT map[string][]string) *TitleTask {
	return &TitleTask{
		Id:        id,
		TaskId:    taskId,
		Type:      Title,
		NextTasks: nT,
	}
}

func (t TitleTask) GetFileName() string {
	// task type didn't has a file assossiated
	return ""
}

func (t TitleTask) CreateNode() *Node {
	nD := TitleNodeData{
		label: t.Name,
	}
	return CreateNode(t.TaskId, t.Type, nD.GetData())
}

func (t TitleTask) CreateEdges() *[]Edge {
	edges := []Edge{}
	for k, vals := range t.NextTasks {
		for _, val := range vals {
			e := Edge{
				Id:     fmt.Sprintf("%s-%s", t.Id, val),
				Source: t.Id,
				Target: val,
				Label:  k,
			}
			edges = append(edges, e)
		}
	}
	return &edges
}

type ConditionTask struct {
	Id        string              `json:"id"`
	NextTasks map[string][]string `json:"next_tasks"`
	Type      PlaybookTaskType    `json:"type"`
	Name      string              `json:"name"`
	TaskId    string              `json:"task_id"`
	// FUTURE maybe add the condition themself to the graph
}

func CreateConditionTask(id, taskId string, nT map[string][]string) *ConditionTask {
	return &ConditionTask{
		Id:        id,
		TaskId:    taskId,
		Type:      Title,
		NextTasks: nT,
	}
}

func (t ConditionTask) GetNextTask() map[string][]string {
	return t.NextTasks
}

func (t ConditionTask) GetFileName() string {
	// task type didn't has a file assossiated
	return ""
}

func (t ConditionTask) CreateNode() *Node {
	cD := ConditionNodeData{
		label:      t.Name,
		conditions: []string{},
	}
	for k := range t.NextTasks {
		cD.conditions = append(cD.conditions, k)
	}
	return CreateNode(t.Id, t.Type, cD.GetData())
}

func (t ConditionTask) CreateEdges() *[]Edge {
	edges := []Edge{}
	for k, vals := range t.NextTasks {
		for _, val := range vals {
			e := Edge{
				Id:     fmt.Sprintf("%s-%s", t.Id, val),
				Source: t.Id,
				Target: val,
				Label:  k,
			}
			edges = append(edges, e)
		}
	}
	return &edges
}

type PlaybookTask struct {
	Id          string              `json:"id"`
	NextTasks   map[string][]string `json:"next_tasks"`
	Type        PlaybookTaskType    `json:"type"`
	Name        string              `json:"name"`
	TaskId      string              `json:"task_id"`
	PlaybookId  string              `json:"playbook_id"`
	Description string              `json:"description"`
	Args        []string
}

func CreatePlaybookTask(id, taskId, pbId, desc, name string, args []string, nT map[string][]string) *PlaybookTask {
	return &PlaybookTask{
		Id:          id,
		TaskId:      taskId,
		Type:        playbook,
		NextTasks:   nT,
		Name:        name,
		Description: desc,
		PlaybookId:  pbId,
		Args:        args,
	}
}

func (t PlaybookTask) GetNextTask() map[string][]string {
	return t.NextTasks
}

func (t PlaybookTask) GetFileName() string {
	return fmt.Sprintf("playbook-%s.yml", strings.ReplaceAll(t.Name, " ", "_"))
}

func (t PlaybookTask) CreateNode() *Node {
	pD := PlaybookNodeData{
		label:       t.Name,
		description: t.Description,
		args:        t.Args,
	}
	return CreateNode(t.TaskId, t.Type, pD.GetData())
}

func (t PlaybookTask) CreateChildNodes() *[]Node {
	return &[]Node{}
}

func (t PlaybookTask) CreateEdges() *[]Edge {
	edges := []Edge{}
	for k, vals := range t.NextTasks {
		for _, val := range vals {
			e := Edge{
				Id:     fmt.Sprintf("%s-%s", t.Id, val),
				Source: t.Id,
				Target: val,
				Label:  k,
			}
			edges = append(edges, e)
		}
	}
	return &edges
}

type AutomationTask struct {
	Id          string              `json:"id"`
	NextTasks   map[string][]string `json:"next_tasks"`
	Type        PlaybookTaskType    `json:"type"`
	Name        string              `json:"name"`
	TaskId      string              `json:"task_id"`
	Description string              `json:"description"`
	Args        []string            `json:"args"`
}

func CreateAutomationTask(id, taskId, desc, name string, args []string, nT map[string][]string) *AutomationTask {
	return &AutomationTask{
		Id:          id,
		TaskId:      taskId,
		Type:        Automation,
		NextTasks:   nT,
		Name:        name,
		Description: desc,
		Args:        args,
	}
}

func (t AutomationTask) GetNextTask() map[string][]string {
	return t.NextTasks
}

func (t AutomationTask) GetFileName() string {
	// BUG - The yaml only contains the uuid from the script not the script name or file name
	return "" /* fmt.Sprintf("automation-%s.yml", strings.ReplaceAll(t.Name, " ", "_")) */
}

func (t AutomationTask) CreateNode() *Node {
	nD := AutomationNodeData{
		label: t.Name,
		args:  []string{},
	}
	nD.args = append(nD.args, t.Args...)
	return CreateNode(t.Id, t.Type, nD.GetData())
}

func (t AutomationTask) CreateEdges() *[]Edge {
	edges := []Edge{}
	for k, vals := range t.NextTasks {
		for _, val := range vals {
			e := Edge{
				Id:     fmt.Sprintf("%s-%s", t.Id, val),
				Source: t.Id,
				Target: val,
				Label:  k,
			}
			edges = append(edges, e)
		}
	}
	return &edges
}

type CollectionTask struct {
	Id          string              `json:"id"`
	NextTasks   map[string][]string `json:"next_tasks"`
	Type        PlaybookTaskType    `json:"type"`
	Name        string              `json:"name"`
	TaskId      string              `json:"task_id"`
	Description string              `json:"description"`
	// TODO -> Could be interessting to see what questions are asked
}

func CreateCollectionTask(id, taskId, name, desc string, nextTasks map[string][]string) *CollectionTask {
	return &CollectionTask{
		Id:          id,
		Type:        collection,
		Name:        name,
		Description: desc,
		TaskId:      taskId,
		NextTasks:   nextTasks,
	}
}

func (t CollectionTask) GetNextTask() map[string][]string {
	return t.NextTasks
}

func (t CollectionTask) GetFileName() string {
	// task type didn't has a file assossiated
	return ""
}

func (t CollectionTask) CreateNode() *Node {
	nD := CollectionNodeData{
		label:       t.Name,
		description: t.Description,
	}
	return CreateNode(t.TaskId, t.Type, nD.GetData())
}

func (t CollectionTask) CreateEdges() *[]Edge {
	edges := []Edge{}
	for k, vals := range t.NextTasks {
		for _, val := range vals {
			e := Edge{
				Id:     fmt.Sprintf("%s-%s", t.Id, val),
				Source: t.Id,
				Target: val,
				Label:  k,
			}
			edges = append(edges, e)
		}
	}
	return &edges
}

type PlaybookTaskType string

const (
	collection PlaybookTaskType = "collection"
	playbook   PlaybookTaskType = "playbook"
	Start      PlaybookTaskType = "start"
	Automation PlaybookTaskType = "automation"
	Condition  PlaybookTaskType = "condition"
	Title      PlaybookTaskType = "title"
)

func toTaskType(str string) PlaybookTaskType {
	str = strings.ToLower(str)
	switch str {
	case "collection":
		return collection
	case "playbook":
		return playbook
	case "start":
		return Start
	case "builtin":
		return Automation
	case "automation":
		return Automation
	case "condition":
		return Condition
	case "title":
		return Title
	}
	return "unknown"
}
