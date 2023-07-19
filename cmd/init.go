package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize repository",
	Long:  "Initialize repository with default values",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initializng repository")
		fmt.Println("git init")
		fmt.Println("creating .gitignore")
		fmt.Println("creating README.md")
		fmt.Printf("creating LICENSE(%s)\n", License)
	},
}
