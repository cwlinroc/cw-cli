package template

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateCode(t *testing.T) {
	dir := t.TempDir()

	if err := GenerateCode(dir, "app"); err != nil {
		t.Fatal(err)
	}

	content, err := os.ReadFile(filepath.Join(dir, "app.code-workspace"))
	if err != nil {
		t.Fatal("app.code-workspace was not created:", err)
	}
	if string(content) != code_Template {
		t.Errorf("got %q, want %q", content, code_Template)
	}
}
