package initialize

import (
	"fmt"
	"log"
	"os"

	"github.com/a-camarillo/quick-repo/cmd/directory"
	"github.com/spf13/cobra"
)

var tsCommand = &cobra.Command{
	Use: "ts",
	Short: "Initialize a Typescript Repository",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello Typescript")
		dir := directory.NewDirectory()
		pwd, err := os.Getwd()

		if err != nil {
			log.Fatal(err)
		}

		dir.Path = pwd
		dir.License = License
		dir.ProjectType = "ts"

		dir.CreateLicense()

		if ReadMe {
			dir.ReadMe = ReadMe
			dir.CreateReadMe()
		}
		
		if Contributing {
			dir.Contribution = Contributing
			dir.CreateContributing()
		}

		err = dir.CreateGitIgnore()
	
		if err != nil {
			log.Fatal(err)
		}
		err = dir.CreateProjectFiles()

		if err != nil {
			log.Fatal(err)
		}
		err = dir.InitializeRepository()

		if err != nil {
			log.Fatal(err)
		}
	},
}
