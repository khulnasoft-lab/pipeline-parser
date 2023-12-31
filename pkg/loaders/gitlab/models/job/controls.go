package job

import (
	"github.com/khulnasoft-lab/pipeline-parser/pkg/consts"
	"github.com/khulnasoft-lab/pipeline-parser/pkg/loaders/utils"
	"github.com/khulnasoft-lab/pipeline-parser/pkg/models"
	"gopkg.in/yaml.v3"
)

// Control represents "job:except/only"
type Controls struct {
	Refs       []string
	Variables  []string
	Changes    []string
	Kubernetes string

	FileReference *models.FileReference
}

func (c *Controls) UnmarshalYAML(node *yaml.Node) error {
	c.FileReference = utils.GetFileReference(node)
	if node.Tag == consts.SequenceTag {
		refs, err := utils.ParseYamlStringSequenceToSlice(node, "Controls")
		if err != nil {
			return err
		}
		c.Refs = refs
		return nil
	}

	return utils.IterateOnMap(node, func(key string, value *yaml.Node) error {
		switch key {
		case "refs":
			refs, err := utils.ParseYamlStringSequenceToSlice(value, "Controls.refs")
			if err != nil {
				return err
			}
			c.Refs = refs
		case "variables":
			variables, err := utils.ParseYamlStringSequenceToSlice(value, "Controls.variables")
			if err != nil {
				return err
			}
			c.Variables = variables
		case "changes":
			changes, err := utils.ParseYamlStringSequenceToSlice(value, "Controls.changes")
			if err != nil {
				return err
			}
			c.Changes = changes
		case "kubernetes":
			c.Kubernetes = value.Value
		}
		return nil
	}, "Controls")
}
