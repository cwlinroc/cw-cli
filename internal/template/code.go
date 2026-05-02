package template

import (
	"fmt"
	"os"
	"path/filepath"
)

func GenerateCode(targetDirect string, fileName string) error {

	file, err := os.Create(filepath.Join(targetDirect, fileName+".code-workspace"))

	if err != nil {
		fmt.Println("Error creating file: ", err)
		return err
	}

	defer file.Close()

	file.WriteString(code_Template)

	return nil
}

var code_Template string = `{
	"folders": [
		{
			"path": "."
		}
	],
	"settings": {}
}`
