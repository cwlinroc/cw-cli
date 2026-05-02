package template

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateEditorConfig(t *testing.T) {
	dir := t.TempDir()

	if err := GenerateEditorConfig(dir); err != nil {
		t.Fatal(err)
	}

	content, err := os.ReadFile(filepath.Join(dir, ".editorconfig"))
	if err != nil {
		t.Fatal(".editorconfig was not created:", err)
	}
	if string(content) != editorConfig_Template {
		t.Errorf("got %q, want %q", content, editorConfig_Template)
	}
}
