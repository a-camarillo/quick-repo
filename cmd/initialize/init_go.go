package initialize

import (
	"fmt"

	"github.com/spf13/cobra"
)


var goCommand = &cobra.Command{
	Use: "go",
	Short: "Initialize a Go Repository",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello Go")
	},
}

