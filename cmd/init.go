package cmd

import (
	"github.com/fatih/color"
	"github.com/goreleaser/goreleaser/internal/static"
	"github.com/goreleaser/goreleaser/pkg/config"
	"github.com/goreleaser/goreleaser/pkg/interactive"
	"gopkg.in/yaml.v2"
	"os"

	"github.com/apex/log"
	"github.com/spf13/cobra"
)

type initCmd struct {
	cmd         *cobra.Command
	config      string
	interactive bool
}

func newInitCmd() *initCmd {
	var root = &initCmd{}
	var cmd = &cobra.Command{
		Use:           "init",
		Aliases:       []string{"i"},
		Short:         "Generates a .goreleaser.yml file",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			f, err := os.OpenFile(root.config, os.O_WRONLY|os.O_CREATE|os.O_TRUNC|os.O_EXCL, 0644)
			if err != nil {
				return err
			}
			defer f.Close()

			if !root.interactive {
				log.Infof(color.New(color.Bold).Sprintf("Generating %s file", root.config))
				if _, err := f.WriteString(static.ExampleConfig); err != nil {
					return err
				}
			} else {
				log.Infof(color.New(color.Bold).Sprint("Entering interactive mode"))
				cb := interactive.NewConfigBuilder()
				var (
					project config.Project
					err     error
				)
				if project, err = cb.Run(); err != nil {
					return err
				}

				if err := yaml.NewEncoder(f).Encode(project); err != nil {
					return err
				}
			}

			log.WithField("file", root.config).Info("config created; please edit accordingly to your needs")
			return nil
		},
	}

	cmd.Flags().StringVarP(&root.config, "config", "f", ".goreleaser.yml", "Load configuration from file")
	cmd.Flags().BoolVarP(&root.interactive, "interactive", "i", false, "Interactively generate goreleaser config file")

	root.cmd = cmd
	return root
}
