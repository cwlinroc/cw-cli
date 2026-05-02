package template

import (
	"fmt"
	"os"
	"path/filepath"
)

func GeneratePage(targetDirect string) error {

	baseDir := filepath.Base(targetDirect)

	file, err := os.Create(filepath.Join(targetDirect, "+page.svelte"))

	if err != nil {
		fmt.Println("Error creating file: ", err)
		return err
	}

	defer file.Close()

	fileText := fmt.Sprintf(sv_page_Template, baseDir, baseDir)

	file.WriteString(fileText)

	return nil
}

var sv_page_Template string = `<svelte:head>
    <title>%s</title>
    <meta name="description" content="%s" />
</svelte:head>`
