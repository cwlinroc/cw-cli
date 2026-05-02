package template

import (
	"fmt"
	"os"
	"path/filepath"
)

func GenerateCode(targetDirect string, fileName string) error {
	err := os.WriteFile(filepath.Join(targetDirect, fileName+".code-workspace"), []byte(code_Template), 0o644)
	if err != nil {
		return fmt.Errorf("write code workspace: %w", err)
	}

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
