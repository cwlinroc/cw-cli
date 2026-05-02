package template

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGeneratePage(t *testing.T) {
	base := t.TempDir()
	dir := filepath.Join(base, "my-page")
	if err := os.Mkdir(dir, 0o755); err != nil {
		t.Fatal(err)
	}

	if err := GeneratePage(dir); err != nil {
		t.Fatal(err)
	}

	content, err := os.ReadFile(filepath.Join(dir, "+page.svelte"))
	if err != nil {
		t.Fatal("+page.svelte was not created:", err)
	}

	body := string(content)
	if !strings.Contains(body, "<title>my-page</title>") {
		t.Errorf("expected <title>my-page</title> in output, got:\n%s", body)
	}
	if !strings.Contains(body, `content="my-page"`) {
		t.Errorf(`expected content="my-page" in output, got:\n%s`, body)
	}
}
