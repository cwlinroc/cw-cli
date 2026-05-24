package template

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateClaude(t *testing.T) {
	dir := t.TempDir()

	if err := GenerateClaude(dir); err != nil {
		t.Fatal(err)
	}

	content, err := os.ReadFile(filepath.Join(dir, "CLAUDE.md"))
	if err != nil {
		t.Fatal("CLAUDE.md was not created:", err)
	}
	if string(content) != claude_Template {
		t.Errorf("got %q, want %q", content, claude_Template)
	}
}
