package common

import (
	gitlabModels "github.com/khulnasoft-labs/pipeline-parser/pkg/loaders/gitlab/models/common"
	"github.com/khulnasoft-labs/pipeline-parser/pkg/models"
	parsersUtils "github.com/khulnasoft-labs/pipeline-parser/pkg/parsers/utils"
	"github.com/khulnasoft-labs/pipeline-parser/pkg/utils"
)

func ParseRunner(image *gitlabModels.Image) *models.Runner {
	if image == nil {
		return nil
	}
	registry, namespace, imageName, tag := parsersUtils.ParseImageName(image.Name)
	if namespace != "" {
		imageName = namespace + "/" + imageName
	}

	return &models.Runner{
		DockerMetadata: &models.DockerMetadata{
			Image:       utils.GetPtrOrNil(imageName),
			Label:       utils.GetPtrOrNil(tag),
			RegistryURL: utils.GetPtrOrNil(registry),
		},
		FileReference: image.FileReference,
	}
}
