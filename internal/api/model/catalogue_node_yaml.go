package model

import "strings"

func (m *CatalogueNode)  SetID(id *int){
	m.ID = int32(*id)
	*id++
	for i := range m.Nodes{
		m.Nodes[i].SetID(id)
	}
}

func (m *CatalogueNode) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var x struct {
		Name string          `yaml:"name,omitempty"`
		Tags  []string        `yaml:"tags,omitempty"`
		Nodes []*CatalogueNode `yaml:"nodes,omitempty"`
	}
	if err := unmarshal(&x); err != nil {
		return err
	}
	if len(x.Name) == 0 && len(x.Tags) > 0 {
		x.Name = x.Tags[0]
		for _,s := range x.Tags[1:] {
			x.Name += ", " + strings.ToLower(s)
		}
	}
	m.Name = x.Name
	m.Tags = x.Tags
	m.Nodes = x.Nodes
	return nil
}

