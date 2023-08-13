package models

import (
	"github.com/khulnasoft-lab/pipeline-parser/pkg/consts"
	loadersUtils "github.com/khulnasoft-lab/pipeline-parser/pkg/loaders/utils"
	"github.com/khulnasoft-lab/pipeline-parser/pkg/models"
	"gopkg.in/yaml.v3"
)

type Pool struct {
	Name          string   `yaml:"name"`
	Demands       []string `yaml:"demands"`
	VmImage       string   `yaml:"vmImage"`
	FileReference *models.FileReference
}

func (p *Pool) UnmarshalYAML(node *yaml.Node) error {
	p.FileReference = loadersUtils.GetFileReference(node)
	if node.Tag == consts.StringTag {
		p.Name = node.Value
		return nil
	}

	p.FileReference.StartRef.Line--
	return loadersUtils.IterateOnMap(node, func(key string, value *yaml.Node) error {
		switch key {
		case "name":
			p.Name = value.Value
		case "demands":
			var demands []string
			if err := loadersUtils.ParseSequenceOrOne(value, &demands); err != nil {
				return err
			}
			p.Demands = demands
		case "vmImage":
			p.VmImage = value.Value
		}

		return nil
	}, "Pool")
}
