package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/huh"
)

func PickDir(theme *huh.Theme) (targetDirect string, err error) {
	currentDir, err := os.Getwd()

	if err != nil {
		return "", errors.New("Error getting current directory " + err.Error())
	}

	dir := currentDir

	for {
		fmt.Println("Directory: ", dir)

		items, err := os.ReadDir(dir)

		if err != nil {
			return "", errors.New("Error reading directory " + err.Error())
		}

		var entryNames []string
		for _, v := range items {
			if v.IsDir() {
				entryNames = append(entryNames, v.Name())
			}

		}
		entryOptions := make([]huh.Option[string], len(entryNames))

		for i, v := range entryNames {
			entryOptions[i] = huh.NewOption(v, v)
		}

		baseDir := filepath.Base(dir)

		var huhOptions []huh.Option[string]

		if baseDir == "" || baseDir == "/" || baseDir == "\\" {
			huhOptions = append(huh.NewOptions("<current directory>", "<mkdir>"), entryOptions...)
		} else {
			huhOptions = append(huh.NewOptions("<current directory>", "<mkdir>", "../"), entryOptions...)
		}

		var result string

		err = huh.NewSelect[string]().
			Title("Select File Path").
			Options(huhOptions...).
			Value(&result).
			WithTheme(theme).
			Run()

		if err != nil {
			return "", errors.New("Error getting directory name " + err.Error())
		}

		println("Result: ", result)

		switch result {

		case "<current directory>":
			return dir, nil

		case "../":
			dir = filepath.Dir(dir)

		case "<mkdir>":
			var newDir string

			err := huh.NewInput().
				Title("New Directory Name").
				Value(&newDir).
				WithTheme(theme).
				Run()

			if err != nil {
				return "", errors.New("Error getting new directory name " + err.Error())
			}

			dir = filepath.Join(dir, newDir)

			err = os.Mkdir(dir, 0755)

			if err != nil {
				return "", errors.New("Error creating directory " + err.Error())
			}

		case "<exit>":
			return "", errors.New("user exited")

		default:
			dir = filepath.Join(dir, result)

		}
	}
}
