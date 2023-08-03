package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/a-camarillo/quick-repo/cmd/initialize"
	"github.com/spf13/cobra"
)

var PWD string

var rootCmd = &cobra.Command{
	Use:     "quick-repo [OPTIONS]",
	Version: "0.0.1",
	Short:   "Quick git repository setup",
	Long:    "Bootstrap git repositories with necessary files like readme, gitignore, license, etc",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func Execute() {
	rootCmd.AddCommand(initialize.InitCommand)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
