package cmd

import (
	"github.com/spf13/cobra"
)

var (
	AppCommand = &cobra.Command{
		Use:   "app",
		Short: "Application's command",
		Long:  "app command support create, deploy, upgrade and destroy an application",
	}

	appParams = struct {
		App         string
		Environment string
		Tag         string
	}{}
)

func init() {
	up := &cobra.Command{
		Use:   "up",
		Short: "deploy or upgrade an exists application",
		Run: func(cmd *cobra.Command, args []string) {
			if appParams.App == "" {
				cmd.Help()
			}
		},
	}
	AppCommand.PersistentFlags().StringVarP(&appParams.App, "app", "a", "", "Application name, eg: nginx")
	AppCommand.PersistentFlags().StringVarP(&appParams.Environment, "environment", "e", "develop", "Environment: develop | stage | prod")
	AppCommand.PersistentFlags().StringVarP(&appParams.Tag, "tag", "t", "latest", "Image tag, eg: 1.2.3")

	AppCommand.AddCommand(up)
}
