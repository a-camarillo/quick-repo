package cmd

var ReadMe bool
var License string
var CodeOfConduct bool
var Contributing bool

func init() {

	rootCmd.Flags().BoolVar(&ReadMe, "readme", true, "Initialize with a README.md markdown file")
	rootCmd.Flags().BoolVar(&CodeOfConduct, "coc", false, "Initialize with a CODE_OF_CONDUCT.md file")
	rootCmd.Flags().BoolVar(&Contributing, "contrib", false, "Initialize with a CONTRIBUTING.md file")
	rootCmd.Flags().StringVarP(&License, "license", "l", "MIT", "Initialize with a LICENSE file")
}
