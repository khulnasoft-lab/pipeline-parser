package job

import (
	"github.com/khulnasoft-lab/pipeline-parser/pkg/consts"
	"github.com/khulnasoft-lab/pipeline-parser/pkg/loaders/utils"
	"gopkg.in/yaml.v3"
)

type Inherit struct {
	Default   *InheritValues `yaml:"default,omitempty"`
	Variables *InheritValues `yaml:"variables,omitempty"`
}

type InheritValues struct {
	Enabled *bool
	Keys    []string `yaml:"default_keys,omitempty"`
}

func (i *InheritValues) UnmarshalYAML(node *yaml.Node) error {
	if node.Tag == consts.BooleanTag {
		i.Enabled = utils.MustParseYamlBooleanValue(node)
		return nil
	}

	if node.Tag == consts.SequenceTag {
		keys, err := utils.ParseYamlStringSequenceToSlice(node, "InheritValues")
		if err != nil {
			return err
		}

		i.Keys = keys
		return nil
	}

	return consts.NewErrInvalidYamlTag(node.Tag, "InheritValues")
}
