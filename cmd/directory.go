package cmd

import (
	"embed"
	"fmt"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
)

//go:embed templates/*
var fs embed.FS

func InitializeRepository(path string) error {
	_, err := git.PlainInit(path, false)
	if err != nil {
		return err
	}
	fmt.Printf("Initializing Repository at %s/.git/\n", path)
	return nil
}

func CreateReadMe() error {
	_, err := os.Create("README.md")
	if err != nil {
		return err
	}
	return nil
}

func CreateLicense(license string) error {
	switch strings.ToLower(license) {
	case "mit":
		f, err := fs.ReadFile("templates/licenses/LICENSE-MIT.txt")
		if err != nil {
			return err
		}
		err = os.WriteFile("LICENSE.txt", f, 0655)
		if err != nil {
			return err
		}
	case "apache":
		f, err := fs.ReadFile("templates/licenses/LICENSE-APACHE.txt")
		if err != nil {
			return err
		}
		err = os.WriteFile("LICENSE.txt", f, 0655)
		if err != nil {
			return err
		}
	case "apgl":
		f, err := fs.ReadFile("templates/licenses/LICENSE-APGL.txt")
		if err != nil {
			return err
		}
		err = os.WriteFile("LICENSE.txt", f, 0655)
		if err != nil {
			return err
		}
	case "gpl":
		f, err := fs.ReadFile("templates/licenses/LICENSE-GPL.txt")
		if err != nil {
			return err
		}
		err = os.WriteFile("LICENSE.txt", f, 0655)
		if err != nil {
			return err
		}
	case "lpgl":
		f, err := fs.ReadFile("templates/licenses/LICENSE-LPGL.txt")
		if err != nil {
			return err
		}
		err = os.WriteFile("LICENSE.txt", f, 0655)
		if err != nil {
			return err
		}
	case "moz":
		f, err := fs.ReadFile("templates/licenses/LICENSE-MOZ.txt")
		if err != nil {
			return err
		}
		err = os.WriteFile("LICENSE.txt", f, 0655)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("%s is not a valid license. Please enter a valid license: MIT, APACHE, APGL, GPL, LPGL, MOZ", license)
	}
	return nil
}

func CreateContributing() error {
	f, err := fs.ReadFile("templates/CONTRIBUTING.md")
	if err != nil {
		return err
	}
	err = os.WriteFile("CONTRIBUTING.md", f, 0655)
	if err != nil {
		return err
	}
	return nil
}

func CreateCodeOfConduct() error {
	f, err := fs.ReadFile("templates/CODE_OF_CONDUCT.md")
	if err != nil {
		return err
	}
	err = os.WriteFile("CODE_OF_CONDUCT.md", f, 0655)
	if err != nil {
		return err
	}
	return nil
}
