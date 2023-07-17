package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "quick-repo [OPTIONS]",
	Version: "0.0.1",
	Short:   "Quick git repository setup",
	Long:    "Bootstrap git repositories with necessary files like readme, gitignore, license, etc.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello World")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
