package initialize

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	InitCommand.AddCommand(goCommand)
	InitCommand.AddCommand(tsCommand)
}

var InitCommand = &cobra.Command{
		Use: "init",
		Short: "Initialize repository.",
		Long: "Initialize repository with default values",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello Init")
		},
}
