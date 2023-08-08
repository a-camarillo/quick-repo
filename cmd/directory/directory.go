package directory

import (
	"embed"
	"fmt"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
)

//go:embed templates/*
var fs embed.FS

// Directory is a type for holding all of the directory information relative
// to quick-repo flags
type Directory struct {
	Path		string
	ReadMe			bool	
	License			string
	Contribution	bool
	CodeOfConduct	bool
	GitIgnore		string
	ProjectType		string
}

func NewDirectory() *Directory {
	return &Directory{}
}

func (d *Directory) InitializeRepository() error {
	_, err := git.PlainInit(d.Path, false)
	if err != nil {
		return err
	}
	fmt.Printf("Initializing Repository at %s/.git/\n", d.Path)
	return nil
}

func (d *Directory) CreateReadMe() error {
	if !d.ReadMe {
		err := fmt.Errorf("specification for creating a README returned false")
		return err
	} else {
		_, err := os.Create("README.md")
		if err != nil {
			return err
		}
		return nil
	}
}

func (d *Directory) CreateLicense() error {
	switch strings.ToLower(d.License) {
	case "mit":
		f, err := fs.ReadFile("templates/licenses/LICENSE-MIT.txt")
		if err != nil {
			return err
		}
		err = os.WriteFile("LICENSE.txt", f, 0644)
		if err != nil {
			return err
		}
	case "apache":
		f, err := fs.ReadFile("templates/licenses/LICENSE-APACHE.txt")
		if err != nil {
			return err
		}
		err = os.WriteFile("LICENSE.txt", f, 0644)
		if err != nil {
			return err
		}
	case "apgl":
		f, err := fs.ReadFile("templates/licenses/LICENSE-APGL.txt")
		if err != nil {
			return err
		}
		err = os.WriteFile("LICENSE.txt", f, 0644)
		if err != nil {
			return err
		}
	case "gpl":
		f, err := fs.ReadFile("templates/licenses/LICENSE-GPL.txt")
		if err != nil {
			return err
		}
		err = os.WriteFile("LICENSE.txt", f, 0644)
		if err != nil {
			return err
		}
	case "lpgl":
		f, err := fs.ReadFile("templates/licenses/LICENSE-LPGL.txt")
		if err != nil {
			return err
		}
		err = os.WriteFile("LICENSE.txt", f, 0644)
		if err != nil {
			return err
		}
	case "moz":
		f, err := fs.ReadFile("templates/licenses/LICENSE-MOZ.txt")
		if err != nil {
			return err
		}
		err = os.WriteFile("LICENSE.txt", f, 0644)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("%s is not a valid license. Please enter a valid license: MIT, APACHE, APGL, GPL, LPGL, MOZ", d.License)
	}
	return nil
}


func (d *Directory) CreateContributing() error {
	if !d.Contribution {
		err := fmt.Errorf("specification for creating a CONTRIBUTION returned false")
		return err
	}	else {
		f, err := fs.ReadFile("templates/CONTRIBUTING.md")
		if err != nil {
			return err
		}
		err = os.WriteFile("CONTRIBUTING.md", f, 0644)
		if err != nil {
			return err
		}
		return nil
	}
}

func (d *Directory) CreateCodeOfConduct() error {
	if !d.CodeOfConduct {
		err := fmt.Errorf("specification for creating a CODE_OF_CONDUCT returned false")
		return err
	} else {
		f, err := fs.ReadFile("templates/CODE_OF_CONDUCT.md")
		if err != nil {
			return err
		}
		err = os.WriteFile("CODE_OF_CONDUCT.md", f, 0644)
		if err != nil {
			return err
		}
		return nil
	}
}

func (d *Directory) CreateGitIgnore() error {
	switch d.ProjectType {
		case "go": 
			f, err := fs.ReadFile("templates/go_templates/gitignore-go")
			if err != nil {
				return err
			}
			err = os.WriteFile(".gitignore", f, 0644)
			if err != nil {
				return err
			}
			return nil 	
		case "ts": 
			f, err := fs.ReadFile("templates/ts_templates/gitignore-ts")
			if err != nil {
				return err
			}
			err = os.WriteFile(".gitignore", f, 0644)
			if err != nil {
				return err
			}
			return nil
		case "py": 
			f, err := fs.ReadFile("templates/py_templates/gitignore-py")
			if err != nil {
				return err
			}
			err = os.WriteFile(".gitignore", f, 0644)
			if err != nil {
				return err
			}
			return nil 	
		default: 
			f, err := fs.ReadFile("templates/.gitignore")
			if err != nil {
				return err
			}
			err = os.WriteFile(".gitignore", f, 0644)
			if err != nil {
				return err
			}
			return nil
	}
}

func (d *Directory) CreateProjectFiles() error {
	switch d.ProjectType {
		case "go":
			main, err := fs.ReadFile("templates/go_templates/main.go")
			if err != nil {
				return err
			}
			mod, err := fs.ReadFile("templates/go_templates/template_go.mod")
			if err != nil {
				return err 
			}	
			err = os.WriteFile("main.go", main, 0644)
			if err != nil {
				return err
			}
			err = os.WriteFile("go.mod", mod, 0644)
			if err != nil {
				return err
			}
			return nil
		case "ts":
			pkg, err := fs.ReadFile("templates/ts_templates/package.json")
			if err != nil {
				return err
			}
			idx, err := fs.ReadFile("templates/ts_templates/index.ts")
			if err != nil {
				return err
			}
			config, err := fs.ReadFile("templates/ts_templates/tsconfig.json")
			if err != nil {
				return err 
			}	
			err = os.WriteFile("package.json", pkg, 0644)
			if err != nil {
				return err
			}
			err = os.WriteFile("tsconfig.json", config, 0644)
			if err != nil {
				return err
			}
			err = os.WriteFile("index.ts", idx, 0644)
			if err != nil {
				return err
			}
			return nil
		case "py":
			main, err := fs.ReadFile("templates/py_templates/main.py")
			if err != nil {
				return err 
			}
			err = os.WriteFile("main.py", main, 0644)
			if err != nil {
				return err
			}
			return nil
		default:
			return nil
	}
}
