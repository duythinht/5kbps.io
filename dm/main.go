package main

import (
	"github.com/spf13/cobra"
	"go.5kbps.io/dm/cmd"
)

func main() {
	dm := cobra.Command{
		Use:   "dm",
		Short: "Simple Helm wrapper Deploy Manager",
		Long:  "This tool just be made for my work, that make deploy helm charts more efficient and avoid human's mistake",
	}

	dm.AddCommand(cmd.AppCommand)

	_ = dm.Execute()
}
