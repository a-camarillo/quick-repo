package initialize

var ReadMe bool
var License string
var CodeOfConduct bool
var Contributing bool

func init() {
	InitCommand.Flags().BoolVar(&ReadMe, "readme", true, "Initialize with a README.md markdown file")
	InitCommand.Flags().BoolVar(&CodeOfConduct, "coc", false, "Initialize with a CODE_OF_CONDUCT.md file")
	InitCommand.Flags().BoolVar(&Contributing, "contrib", false, "Initialize with a CONTRIBUTING.md file")
	InitCommand.Flags().StringVarP(&License, "license", "l", "MIT", "Initialize with a LICENSE file")
}
