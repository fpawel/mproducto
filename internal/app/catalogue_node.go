package app

import (
	"github.com/powerman/must"
	"gopkg.in/yaml.v2"
	"strings"
)

func ReadCatalogueNodes() (result []CatalogueNode) {
	b := must.ReadFile("catalogue.yml")
	must.AbortIf(yaml.Unmarshal(b, &result))
	return
}

type CatalogueNode struct {
	Title string          `json:"title"`
	Tags  []string        `json:"tags"`
	Nodes []CatalogueNode `json:"nodes"`
}

func (u *CatalogueNode) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var x struct {
		Title string          `yaml:"title,omitempty"`
		Tags  []string        `yaml:"tags,omitempty"`
		Nodes []CatalogueNode `yaml:"nodes,omitempty"`
	}
	if err := unmarshal(&x); err != nil {
		return err
	}
	if len(x.Title) == 0 && len(x.Tags) > 0 {
		x.Title = strings.Title(x.Tags[0])
		tags := append([]string{}, x.Tags...)
		for i := range tags[1:] {
			x.Title += ", " + strings.ToLower(tags[i])
		}
	}
	u.Title = x.Title
	u.Tags = x.Tags
	u.Nodes = x.Nodes
	return nil
}
