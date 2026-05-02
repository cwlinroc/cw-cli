package template

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"
)

func GenerateCs(targetDirect string, fileName string) error {

	nameSpace, err := cs_namespace(targetDirect)

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

func cs_namespace(targetDirect string) (nameSpace string, err error) {
	var names []string
	dir := targetDirect

OuterLoop:
	for i := 0; i < 50; i++ {
		items, err := os.ReadDir(dir)

		if err != nil {
			return "", errors.New("Error reading directory " + err.Error())
		}

		for _, item := range items {
			if path.Ext(item.Name()) == ".csproj" {
				names = append(names, strings.TrimSuffix(item.Name(), ".csproj"))
				break OuterLoop
			}
		}

		baseDir := filepath.Base(dir)

		if baseDir == "" || baseDir == "/" || baseDir == "\\" {
			break
		}

		names = append(names, baseDir)
		dir = filepath.Dir(dir)
	}

	// for i, j := 0, len(names)-1; i < j; i, j = i+1, j-1 {
	// 	names[i], names[j] = names[j], names[i]
	// }

	slices.Reverse(names)
	return strings.Join(names, "."), nil

}
