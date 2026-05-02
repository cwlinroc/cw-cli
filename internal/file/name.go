package file

import (
	"errors"
	"fmt"
	"strings"

	"github.com/charmbracelet/huh"
)

func GetName(args []string, theme *huh.Theme) (string, error) {
	var fileName string
	var err error

	if len(args) > 0 && args[0] != "" {
		fileName = args[0]
	} else {
		fileName, err = promptName(theme)

		if err != nil {
			return "", errors.New("Error getting file name: " + err.Error())
		}
	}

	fileName = strings.TrimSpace(fileName)

	if fileName == "" {
		return "", errors.New("File name cannot be empty")
	}

	return fileName, nil
}

func promptName(theme *huh.Theme) (fileName string, err error) {

	var _fileName string

	// get file name with huh prompt
	{
		for _fileName == "" {

			err := huh.NewInput().
				Title("File Name").
				Value(&_fileName).
				WithTheme(theme).
				Run()

			if err != nil {
				return "", errors.New("Error getting file name " + err.Error())
			}

			_fileName = strings.TrimSpace(_fileName)

			if _fileName == "" {
				fmt.Println("File name cannot be empty")
			}
		}
	}

	return _fileName, nil
}
