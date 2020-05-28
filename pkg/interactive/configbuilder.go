package interactive

import (
	"bufio"
	"fmt"
	"github.com/goreleaser/goreleaser/pkg/config"
	"os"
)

type ConfigBuilder interface {
	Run() (config.Project, error)
}

//type configOption interface {
//	Apply(project config.Project) error
//}

type configOption func(project *config.Project) error

type configBuilder struct {
	options []configOption
}

func (cb *configBuilder) Run() (config.Project, error) {
	var (
		project config.Project
		err     error
	)

	for _, opt := range cb.options {
		if err = opt(&project); err != nil {
			return project, err
		}
	}
	return project, nil
}

func NewConfigBuilder() ConfigBuilder {
	return &configBuilder{
		options: []configOption{
			projectName,
		},
	}
}

func projectName(project *config.Project) error {
	fmt.Print("Project name: ")
	r := bufio.NewReader(os.Stdin)
	line, _, err := r.ReadLine()
	if err != nil {
		return err
	}
	project.ProjectName = string(line)
	return nil
}
