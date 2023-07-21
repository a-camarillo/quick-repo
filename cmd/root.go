package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var PWD string

var rootCmd = &cobra.Command{
	Use:     "quick-repo [OPTIONS]",
	Version: "0.0.1",
	Short:   "Quick git repository setup",
	Long:    "Bootstrap git repositories with necessary files like readme, gitignore, license, etc",
	Run: func(cmd *cobra.Command, args []string) {
		PWD, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		err = CreateLicense(License)
		if err != nil {
			log.Fatal(err)
		}

		if ReadMe {
			err = CreateReadMe()
			if err != nil {
				log.Fatal(err)
			}
		}

		if CodeOfConduct {
			err = CreateCodeOfConduct()
			if err != nil {
				log.Fatal(err)
			}
		}

		if Contributing {
			err = CreateContributing()
			if err != nil {
				log.Fatal(err)
			}
		}

		err = InitializeRepository(PWD)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
