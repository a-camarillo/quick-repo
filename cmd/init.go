package cmd

import (
	"log"
	"os"

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
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		err = InitializeRepository(pwd)
		if err != nil {
			log.Fatal(err)
		}
		err = CreateReadMe()
		if err != nil {
			log.Fatal(err)
		}
		err = CreateLicense(License)
		if err != nil {
			log.Fatal(err)
		}
	},
}
