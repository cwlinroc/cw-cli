package template

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateGitignore(t *testing.T) {
	t.Run("all languages selected by default", func(t *testing.T) {
		dir := t.TempDir()
		opts := GitignoreOptions{
			Dotnet: true,
			Go:     true,
			Java:   true,
			JsTs:   true,
			Python: true,
		}

		err := GenerateGitignore(dir, opts)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		contentBytes, err := os.ReadFile(filepath.Join(dir, ".gitignore"))
		if err != nil {
			t.Fatalf("failed to read .gitignore: %v", err)
		}
		content := string(contentBytes)

		// Verify headers/sections exist
		expectedSections := []string{
			"Universal",
			"docs/draft/",
			"settings.local.json",
			"Dotnet",
			"Go",
			"Java",
			"Node / JS / TS",
			"Python",
		}
		for _, section := range expectedSections {
			if !strings.Contains(content, section) {
				t.Errorf("expected section %q not found in .gitignore", section)
			}
		}
	})

	t.Run("only specific languages", func(t *testing.T) {
		dir := t.TempDir()
		opts := GitignoreOptions{
			Go:     true,
			Python: true,
		}

		err := GenerateGitignore(dir, opts)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		contentBytes, err := os.ReadFile(filepath.Join(dir, ".gitignore"))
		if err != nil {
			t.Fatalf("failed to read .gitignore: %v", err)
		}
		content := string(contentBytes)

		// Verify selected languages are present
		if !strings.Contains(content, "Go") {
			t.Error("expected Go section to be present")
		}
		if !strings.Contains(content, "Python") {
			t.Error("expected Python section to be present")
		}

		// Verify universal patterns are present
		if !strings.Contains(content, "Universal") {
			t.Error("expected Universal section to be present")
		}
		if !strings.Contains(content, "docs/draft/") {
			t.Error("expected docs/draft/ pattern to be present")
		}
		if !strings.Contains(content, "settings.local.json") {
			t.Error("expected settings.local.json pattern to be present")
		}

		// Verify unselected languages are absent
		if strings.Contains(content, "Dotnet") {
			t.Error("unexpected Dotnet section in .gitignore")
		}
		if strings.Contains(content, "Java") {
			t.Error("unexpected Java section in .gitignore")
		}
		if strings.Contains(content, "Node / JS / TS") {
			t.Error("unexpected Node / JS / TS section in .gitignore")
		}
	})
}
