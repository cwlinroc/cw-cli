package template

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateRazor(t *testing.T) {
	dir := t.TempDir()
	touch(t, filepath.Join(dir, "Web.csproj"))

	if err := GenerateRazor(dir, "Index"); err != nil {
		t.Fatal(err)
	}

	t.Run("creates .cshtml file", func(t *testing.T) {
		if _, err := os.Stat(filepath.Join(dir, "Index.cshtml")); err != nil {
			t.Error("Index.cshtml was not created")
		}
	})

	t.Run("creates .cshtml.cs file", func(t *testing.T) {
		content, err := os.ReadFile(filepath.Join(dir, "Index.cshtml.cs"))
		if err != nil {
			t.Fatal("Index.cshtml.cs was not created:", err)
		}
		if !strings.Contains(string(content), "namespace Web") {
			t.Errorf("expected namespace Web in code-behind, got:\n%s", content)
		}
		if !strings.Contains(string(content), "IndexModel") {
			t.Errorf("expected IndexModel in code-behind, got:\n%s", content)
		}
	})
}
