package template

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateCs(t *testing.T) {
	t.Run("creates file with correct namespace and class name", func(t *testing.T) {
		dir := t.TempDir()
		touch(t, filepath.Join(dir, "My.csproj"))

		if err := GenerateCs(dir, "Foo"); err != nil {
			t.Fatal(err)
		}

		content, err := os.ReadFile(filepath.Join(dir, "Foo.cs"))
		if err != nil {
			t.Fatal(err)
		}
		want := fmt.Sprintf(cs_Template, "My", "Foo")
		if string(content) != want {
			t.Errorf("got %q, want %q", content, want)
		}
	})

	t.Run("strips .cs suffix from file name", func(t *testing.T) {
		dir := t.TempDir()
		touch(t, filepath.Join(dir, "App.csproj"))

		if err := GenerateCs(dir, "Bar.cs"); err != nil {
			t.Fatal(err)
		}

		if _, err := os.Stat(filepath.Join(dir, "Bar.cs")); err != nil {
			t.Error("Bar.cs was not created")
		}
		if _, err := os.Stat(filepath.Join(dir, "Bar.cs.cs")); err == nil {
			t.Error("Bar.cs.cs should not exist")
		}
	})

	t.Run("empty name returns error", func(t *testing.T) {
		dir := t.TempDir()
		err := GenerateCs(dir, "")
		if err == nil {
			t.Error("expected error for empty file name")
		}
	})
}

func touch(t *testing.T, path string) {
	t.Helper()
	if err := os.WriteFile(path, []byte{}, 0o644); err != nil {
		t.Fatal(err)
	}
}
