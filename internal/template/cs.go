package template

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"cw/internal/file"
)

func GenerateCs(targetDirect string, fileName string) error {
	if len(fileName) > 3 && fileName[len(fileName)-3:] == ".cs" {
		fileName = fileName[:len(fileName)-3]
	}

	if fileName == "" {
		return errors.New("file name cannot be empty")
	}

	nameSpace, err := file.ExtractCSNamespace(targetDirect)
	if err != nil {
		return fmt.Errorf("get namespace: %w", err)
	}

	fileText := fmt.Sprintf(cs_Template, nameSpace, fileName)

	err = os.WriteFile(filepath.Join(targetDirect, fileName+".cs"), []byte(fileText), 0o644)
	if err != nil {
		return fmt.Errorf("write class file: %w", err)
	}

	return nil
}

var cs_Template string = `namespace %s;
public class %s
{

}`
