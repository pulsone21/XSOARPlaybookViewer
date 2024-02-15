package dataloader

type GraphContent struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}
type Node struct {
	Id       string            `json:"id"`
	NodeType PlaybookTaskType  `json:"type"`
	Data     map[string]string `json:"data"`
	Position NodePosition      `json:"position"`
}

type NodePosition struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type ChildNode struct {
	Id         string            `json:"id"`
	NodeType   PlaybookTaskType  `json:"type"`
	Data       map[string]string `json:"data"`
	Position   NodePosition      `json:"position"`
	ParentNode string            `json:"parentNode"`
	Extent     string            `json:"extent"`
}

type Edge struct {
	Id     string `json:"id"`
	Source string `json:"source"`
	Target string `json:"target"`
}

// TODO Should add ways to populate the type specfic data layouts since they are only maps.
// Maybe a interface is a better suite here
func CreateNode(id string, nodeType PlaybookTaskType, data map[string]string) *Node {
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

func CreateChildeNode(id, parentId string, nodeType PlaybookTaskType, data map[string]string) *ChildNode {
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
