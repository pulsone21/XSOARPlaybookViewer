package dataloader

type GraphContent struct {
	PlaybookName string  `json:"playbook_name"`
	Nodes        []Nodes `json:"nodes"`
	Edges        []Edge  `json:"edges"`
}

type SubGraph struct {
	Nodes []Nodes
	Edges []Edge
}

type Nodes interface {
	GetId() string
}

type Node struct {
	Id       string                 `json:"id"`
	NodeType PlaybookTaskType       `json:"type"`
	Data     map[string]interface{} `json:"data"`
	Position NodePosition           `json:"position"`
}

func (n Node) GetId() string { return n.Id }

type ChildNode struct {
	Id         string                 `json:"id"`
	NodeType   PlaybookTaskType       `json:"type"`
	Data       map[string]interface{} `json:"data"`
	Position   NodePosition           `json:"position"`
	ParentNode string                 `json:"parentNode"`
	Extent     string                 `json:"extent"`
}

func (n ChildNode) GetId() string { return n.Id }

type Edge struct {
	Id     string `json:"id"`
	Source string `json:"source"`
	Target string `json:"target"`
	Label  string `json:"label"`
}

type NodeData interface {
	GetData() map[string]interface{}
}

type NodePosition struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type ConditionNodeData struct {
	label      string
	conditions []string
}

func (d ConditionNodeData) GetData() map[string]interface{} {
	m := make(map[string]interface{})
	m["label"] = d.label
	m["conditions"] = d.conditions
	return m
}

type TitleNodeData struct {
	label string
}

func (d TitleNodeData) GetData() map[string]interface{} {
	m := make(map[string]interface{})
	m["label"] = d.label
	return m
}

type AutomationNodeData struct {
	label string
	args  []string
}

func (d AutomationNodeData) GetData() map[string]interface{} {
	m := make(map[string]interface{})
	m["label"] = d.label
	m["args"] = d.args
	return m
}

type CollectionNodeData struct {
	label       string
	description string
}

func (d CollectionNodeData) GetData() map[string]interface{} {
	m := make(map[string]interface{})
	m["label"] = d.label
	m["description"] = d.description
	return m
}

type PlaybookNodeData struct {
	label       string
	description string
	args        []string
}

func (d PlaybookNodeData) GetData() map[string]interface{} {
	m := make(map[string]interface{})
	m["label"] = d.label
	m["description"] = d.description
	m["args"] = d.args
	return m
}

// TODO Should add ways to populate the type specfic data layouts since they are only maps.
// Maybe a interface is a better suite here
func CreateNode(id string, nodeType PlaybookTaskType, data map[string]interface{}) *Node {
	return &Node{
		Id:       id,
		NodeType: nodeType,
		Data:     data,
		Position: NodePosition{
			X: 0,
			Y: 0,
		},
	}
}

func CreateChildeNode(id, parentId string, nodeType PlaybookTaskType, data map[string]interface{}) *ChildNode {
	return &ChildNode{
		Id:       id,
		NodeType: nodeType,
		Data:     data,
		Position: NodePosition{
			X: 0,
			Y: 0,
		},
		ParentNode: parentId,
		Extent:     "parent",
	}
}
