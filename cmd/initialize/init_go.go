package initialize

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/a-camarillo/quick-repo/cmd/directory"
)

var goCommand = &cobra.Command{
	Use: "go",
	Short: "Initialize a Go Repository",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello World From Go")
		dir := directory.NewDirectory()
		fmt.Println(dir)
	},
}

