package cmd

import (
	"embed"
	"fmt"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
)

//go:embed templates/licenses/*
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
		err = os.WriteFile("LICENSE.txt", f, 0755)
		if err != nil {
			return err
		}
	case "apache":
		f, err := fs.ReadFile("templates/licenses/LICENSE-APACHE.txt")
		if err != nil {
			return err
		}
		err = os.WriteFile("LICENSE.txt", f, 0755)
		if err != nil {
			return err
		}
	case "apgl":
		f, err := fs.ReadFile("templates/licenses/LICENSE-APGL.txt")
		if err != nil {
			return err
		}
		err = os.WriteFile("LICENSE.txt", f, 0755)
		if err != nil {
			return err
		}
	case "gpl":
		f, err := fs.ReadFile("templates/licenses/LICENSE-GPL.txt")
		if err != nil {
			return err
		}
		err = os.WriteFile("LICENSE.txt", f, 0755)
		if err != nil {
			return err
		}
	case "lpgl":
		f, err := fs.ReadFile("templates/licenses/LICENSE-LPGL.txt")
		if err != nil {
			return err
		}
		err = os.WriteFile("LICENSE.txt", f, 0755)
		if err != nil {
			return err
		}
	case "moz":
		f, err := fs.ReadFile("templates/licenses/LICENSE-MOZ.txt")
		if err != nil {
			return err
		}
		err = os.WriteFile("LICENSE.txt", f, 0755)
		if err != nil {
			return err
		}
	default:
		return nil
	}
	return nil
}
