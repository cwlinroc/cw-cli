package template

import (
	"cw/internal/file"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func GenerateCs(targetDirect string, fileName string) error {

	if len(fileName) > 3 && fileName[len(fileName)-3:] == ".cs" {
		fileName = fileName[:len(fileName)-3]
	}

	if fileName == "" {
		return errors.New("File name cannot be empty")
	}

	nameSpace, err := file.ExtractCSNamespace(targetDirect)

	if err != nil {
		fmt.Println("Error getting namespace: ", err)
		return err
	}

	file, err := os.Create(filepath.Join(targetDirect, fileName+".cs"))

	if err != nil {
		fmt.Println("Error creating file: ", err)
		return err
	}

	defer file.Close()

	fileText := fmt.Sprintf(cs_Template, nameSpace, fileName)

	file.WriteString(fileText)

	return nil
}

var cs_Template string = `namespace %s;
public class %s
{

}`
