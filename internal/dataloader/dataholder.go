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

func (d *DataHolder) GetPlaybookView(id string) error {
}
