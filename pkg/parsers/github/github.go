package github

import (
	githubModels "github.com/khulnasoft-lab/pipeline-parser/pkg/loaders/github/models"
	"github.com/khulnasoft-lab/pipeline-parser/pkg/models"
)

type GitHubParser struct{}

func (g *GitHubParser) Parse(workflow *githubModels.Workflow) (*models.Pipeline, error) {
	var err error
	pipeline := &models.Pipeline{
		Name: &workflow.Name,
	}

	pipeline.Triggers = parseWorkflowTriggers(workflow)

	if workflow.Jobs != nil {
		if pipeline.Jobs, err = parseWorkflowJobs(workflow); err != nil {
			return nil, err
		}
	}

	if pipeline.Defaults, err = parseWorkflowDefaults(workflow); err != nil {
		return nil, err
	}

	return pipeline, nil
}

func parseWorkflowDefaults(workflow *githubModels.Workflow) (*models.Defaults, error) {
	if workflow.Permissions == nil && workflow.Env == nil {
		return nil, nil
	}

	defaults := &models.Defaults{}
	permissions, err := parseTokenPermissions(workflow.Permissions)
	if err != nil {
		return nil, err
	}
	defaults.TokenPermissions = permissions
	defaults.EnvironmentVariables = parseEnvironmentVariablesRef(workflow.Env)

	return defaults, nil
}
