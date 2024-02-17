package dataloader

type DataHolder struct {
	Playbooks map[string]Playbook
}

func (d *DataHolder) GetPlaybooks() *[]Playbook {
	var pbs []Playbook
	for _, pb := range d.Playbooks {
		pbs = append(pbs, pb)
	}
	return &pbs
}

func (d *DataHolder) GetPlaybook(id string) *Playbook {
	pb := d.Playbooks[id]
	return &pb
}

func (d *DataHolder) GetPlaybookView(id string) *GraphContent {
	pb := d.Playbooks[id]
	if pb.Id == "" {
		return nil
	}
	return d.generateGraphContent(&pb)
}

func (d *DataHolder) generateGraphContent(pb *Playbook) *GraphContent {
	gC := GraphContent{
		PlaybookName: pb.Name,
	}
	gC.Nodes, gC.Edges = d.extractNodesEdges(pb)
	return &gC
}

func (d *DataHolder) extractNodesEdges(pb *Playbook) ([]Nodes, []Edge) {
	es := []Edge{}
	ns := []Nodes{}

	for _, t := range pb.Tasks {
		n := t.CreateNode()
		ns = append(ns, *n)
		if t.GetNextTask() != nil {
			es = append(es, *t.CreateEdges()...)
		}

		if n.NodeType == playbook {
			subPb := d.Playbooks[t.GetFileName()]
			cNodes, cEdges := d.extractNodesEdges(&subPb)
			es = append(es, cEdges...)
			ns = append(ns, cNodes...)
		}
	}
	return ns, es
}
