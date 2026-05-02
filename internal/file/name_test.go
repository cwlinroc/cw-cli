package file

import (
	"testing"
)

func TestGetName(t *testing.T) {
	t.Run("returns arg unchanged", func(t *testing.T) {
		got, err := GetName([]string{"Foo"}, nil)
		if err != nil {
			t.Fatal(err)
		}
		if got != "Foo" {
			t.Errorf("got %q, want %q", got, "Foo")
		}
	})

	t.Run("trims surrounding whitespace", func(t *testing.T) {
		got, err := GetName([]string{"  Bar  "}, nil)
		if err != nil {
			t.Fatal(err)
		}
		if got != "Bar" {
			t.Errorf("got %q, want %q", got, "Bar")
		}
	})

	t.Run("whitespace-only arg returns error", func(t *testing.T) {
		_, err := GetName([]string{"   "}, nil)
		if err == nil {
			t.Error("expected error for whitespace-only arg, got nil")
		}
	})
}
