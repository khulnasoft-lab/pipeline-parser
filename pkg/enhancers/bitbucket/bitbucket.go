package bitbucket

import (
	"github.com/khulnasoft-lab/pipeline-parser/pkg/enhancers"
	"github.com/khulnasoft-lab/pipeline-parser/pkg/models"
)

type BitbucketEnhancer struct{}

func (b *BitbucketEnhancer) LoadImportedPipelines(data *models.Pipeline, credentials *models.Credentials, _, _ *string) ([]*enhancers.ImportedPipeline, error) {
	return nil, nil
}

func (b *BitbucketEnhancer) Enhance(data *models.Pipeline, importedPipelines []*enhancers.ImportedPipeline) (*models.Pipeline, error) {
	return data, nil
}

func (b *BitbucketEnhancer) InheritParentPipelineData(parent, child *models.Pipeline) *models.Pipeline {
	return child
}
