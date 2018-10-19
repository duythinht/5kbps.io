package main

import (
	"os"

	"github.com/spf13/cobra"
	"go.5kbps.io/dm/cmd"
)

func main() {
	dm := cobra.Command{
		Use:   "dm",
		Short: "Simple Helm wrapper Deploy Manager",
		Long:  "This tool just be made for my work, that make deploy helm charts more efficient and avoid human's mistake",
	}

	dm.Version = "0.1.0-alpha"

	completion := &cobra.Command{
		Use:   "completion",
		Short: "Output shell completion code",
		Run: func(cmd *cobra.Command, args []string) {
			dm.GenBashCompletion(os.Stdout)
		},
	}

	dm.AddCommand(completion, cmd.AppCommand, cmd.VaultCommand)

	_ = dm.Execute()
}
