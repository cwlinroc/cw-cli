package file

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractCSNamespace(t *testing.T) {
	t.Run("csproj in same directory", func(t *testing.T) {
		dir := t.TempDir()
		touch(t, filepath.Join(dir, "My.csproj"))

		ns, err := ExtractCSNamespace(dir)
		if err != nil {
			t.Fatal(err)
		}
		if ns != "My" {
			t.Errorf("got %q, want %q", ns, "My")
		}
	})

	t.Run("csproj in parent directory", func(t *testing.T) {
		parent := t.TempDir()
		child := filepath.Join(parent, "Controllers")
		if err := os.Mkdir(child, 0o755); err != nil {
			t.Fatal(err)
		}
		touch(t, filepath.Join(parent, "App.csproj"))

		ns, err := ExtractCSNamespace(child)
		if err != nil {
			t.Fatal(err)
		}
		if ns != "App.Controllers" {
			t.Errorf("got %q, want %q", ns, "App.Controllers")
		}
	})

	t.Run("csproj two levels up", func(t *testing.T) {
		root := t.TempDir()
		mid := filepath.Join(root, "Services")
		leaf := filepath.Join(mid, "Auth")
		if err := os.MkdirAll(leaf, 0o755); err != nil {
			t.Fatal(err)
		}
		touch(t, filepath.Join(root, "Core.csproj"))

		ns, err := ExtractCSNamespace(leaf)
		if err != nil {
			t.Fatal(err)
		}
		if ns != "Core.Services.Auth" {
			t.Errorf("got %q, want %q", ns, "Core.Services.Auth")
		}
	})

	t.Run("no csproj returns no error", func(t *testing.T) {
		dir := t.TempDir()
		_, err := ExtractCSNamespace(dir)
		if err != nil {
			t.Errorf("expected no error when no csproj found, got: %v", err)
		}
	})
}

func touch(t *testing.T, path string) {
	t.Helper()
	if err := os.WriteFile(path, []byte{}, 0o644); err != nil {
		t.Fatal(err)
	}
}
