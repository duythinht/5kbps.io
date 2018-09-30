package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.5kbps.io/dm/helm"
)

var (
	AppCommand = &cobra.Command{
		Use: "app",
		PreRun: func(cmd *cobra.Command, args []string) {
		},
		Short: "Application's command",
		Long:  "app command support create, deploy, upgrade and destroy an application",
	}
)

func init() {

	var (
		app string
		env string
		tag string
		key string
	)

	up := &cobra.Command{
		Use:   "up",
		Short: "deploy or upgrade an exists application",
		Run: func(cmd *cobra.Command, args []string) {

			valuesPath := fmt.Sprintf("./app/%s/%s.yaml", app, env)
			set := fmt.Sprintf("--set=image.tag=%s", tag)

			if key == "" {
				if err := helm.Run(append([]string{"upgrade", "-i", app, "-f", valuesPath, "_templates/app", set}, args...)...); err != nil {
					os.Exit(1)
				}
				return
			}

			if err := helm.DeployEncryptedValues(valuesPath, key, append([]string{"upgrade", "-i", app, "_templates/app", set}, args...)...); err != nil {
				os.Exit(1)
			}
		},
	}

	up.Flags().StringVarP(&app, "app", "a", "", "Application name, eg: nginx")
	up.Flags().StringVarP(&env, "env", "e", "develop", "Environment: develop | stage | prod")
	up.Flags().StringVarP(&tag, "tag", "t", "latest", "Image tag, eg: 1.2.3")
	up.Flags().StringVarP(&key, "key", "k", "", "Encryption key")
	up.MarkFlagRequired("app")

	down := &cobra.Command{
		Use:   "down",
		Short: "Destroy an app",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	down.Flags().StringVarP(&app, "app", "a", "", "Application name, eg: nginx")
	down.MarkFlagRequired("app")

	AppCommand.AddCommand(up, down)
}
