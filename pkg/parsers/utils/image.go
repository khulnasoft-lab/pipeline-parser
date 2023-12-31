package utils

import (
	"strings"

	"github.com/khulnasoft-lab/pipeline-parser/pkg/consts"
	"github.com/khulnasoft-lab/pipeline-parser/pkg/models"
	"github.com/khulnasoft-lab/pipeline-parser/pkg/utils"
)

func ParseImageName(imageName string) (string, string, string, string) {
	var registry, namespace, tag string
	image := imageName

	if split := strings.Split(imageName, "/"); len(split) == 3 { // imageName contains registry/repository/image
		registry = split[0]
		namespace = split[1]
		image = split[2]
	} else if len(split) == 2 { // imageName contains repository/image
		namespace = split[0]
		image = split[1]
	}

	if split := strings.Split(image, ":"); len(split) == 2 { // image contains image:tag
		image = split[0]
		tag = split[1]
	}

	return registry, namespace, image, tag
}

func ParseRunnerTag(tag string, runner *models.Runner) *models.Runner {
	if runner == nil {
		return runner
	}

	if tag == consts.SelfHosted {
		runner.SelfHosted = utils.GetPtr(true)
	}

	for os, keywords := range consts.OsToKeywords {
		didFind := false
		for _, keyword := range keywords {
			if strings.Contains(strings.ToLower(tag), keyword) {
				runner.OS = utils.GetPtr(string(os))
				didFind = true
				break
			}
		}
		if didFind {
			break
		}
	}

	for _, arch := range consts.ArchKeywords {
		if strings.Contains(strings.ToLower(tag), arch) {
			runner.Arch = utils.GetPtr(arch)
			break
		}
	}

	return runner
}
