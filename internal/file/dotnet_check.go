package file

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"
)

func ExtractCSNamespace(targetDirect string) (nameSpace string, err error) {
	var names []string
	dir := targetDirect

OuterLoop:
	for range 50 {
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
