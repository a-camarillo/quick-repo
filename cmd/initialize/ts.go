package initialize

import (
	"fmt"

	"github.com/spf13/cobra"
)

var tsCommand = &cobra.Command{
	Use: "ts",
	Short: "Initialize a Typescript Repository",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello Typescript")
	},
}
