package common

import (
	"github.com/khulnasoft-lab/pipeline-parser/pkg/consts"
	"github.com/khulnasoft-lab/pipeline-parser/pkg/loaders/utils"
	"github.com/khulnasoft-lab/pipeline-parser/pkg/models"
	"gopkg.in/yaml.v3"
)

type Image struct {
	Name          string   `yaml:"name"`
	Entrypoint    []string `yaml:"entrypoint"`
	FileReference *models.FileReference
}

func (im *Image) UnmarshalYAML(node *yaml.Node) error {

	im.FileReference = utils.GetFileReference(node)

	if node.Tag == consts.StringTag { // format - "image: image:tag"
		im.Name = node.Value
		im.FileReference.EndRef.Column += len("image: ")
		return nil
	}

	im.FileReference.StartRef.Line--

	return utils.IterateOnMap(node, func(key string, value *yaml.Node) error {
		switch key {
		case "name":
			im.Name = value.Value
		case "entrypoint":
			entrypoints, err := utils.ParseYamlStringSequenceToSlice(value, "Image.entrypoint")
			if err != nil {
				return err
			}
			im.Entrypoint = entrypoints
		}
		return nil
	}, "Image")
}
