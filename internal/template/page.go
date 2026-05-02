package template

import (
	"fmt"
	"os"
	"path/filepath"
)

func GeneratePage(targetDirect string) error {
	baseDir := filepath.Base(targetDirect)

	fileText := fmt.Sprintf(sv_page_Template, baseDir, baseDir)

	err := os.WriteFile(filepath.Join(targetDirect, "+page.svelte"), []byte(fileText), 0o644)
	if err != nil {
		return fmt.Errorf("write page file: %w", err)
	}

	return nil
}

var sv_page_Template string = `<svelte:head>
    <title>%s</title>
    <meta name="description" content="%s" />
</svelte:head>`
