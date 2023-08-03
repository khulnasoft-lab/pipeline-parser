package bitbucket

import (
	"github.com/khulnasoft-labs/pipeline-parser/pkg/loaders/bitbucket/models"
	"gopkg.in/yaml.v3"
)

type BitbucketLoader struct{}

func (b *BitbucketLoader) Load(data []byte) (*models.Pipeline, error) {
	pipeline := &models.Pipeline{}
	err := yaml.Unmarshal(data, pipeline)
	return pipeline, err
}
