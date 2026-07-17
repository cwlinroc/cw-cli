package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestConvertEncodingBOMPolicy(t *testing.T) {
	tests := []struct {
		name    string
		content []byte
		policy  bomPolicy
		want    []byte
	}{
		{
			name:    "add BOM to UTF-8",
			content: []byte("hello"),
			policy:  addBOM,
			want:    append(append([]byte{}, utf8BOM...), []byte("hello")...),
		},
		{
			name:    "add BOM to empty file",
			content: []byte{},
			policy:  addBOM,
			want:    append([]byte{}, utf8BOM...),
		},
		{
			name:    "do not duplicate BOM",
			content: append(append([]byte{}, utf8BOM...), []byte("hello")...),
			policy:  addBOM,
			want:    append(append([]byte{}, utf8BOM...), []byte("hello")...),
		},
		{
			name:    "remove BOM for utf8 command",
			content: append(append([]byte{}, utf8BOM...), []byte("hello")...),
			policy:  removeBOM,
			want:    []byte("hello"),
		},
		{
			name:    "preserve BOM when accepted",
			content: append(append([]byte{}, utf8BOM...), []byte("hello")...),
			policy:  preserveBOM,
			want:    append(append([]byte{}, utf8BOM...), []byte("hello")...),
		},
		{
			name:    "convert Big5 and add BOM",
			content: []byte{0xA4, 0xA4, 0xA4, 0xE5},
			policy:  addBOM,
			want:    append(append([]byte{}, utf8BOM...), []byte("中文")...),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "input.txt")
			if err := os.WriteFile(path, tt.content, 0o600); err != nil {
				t.Fatalf("write test file: %v", err)
			}

			if err := convertEncoding(path, false, tt.policy); err != nil {
				t.Fatalf("convertEncoding: %v", err)
			}

			got, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("read converted file: %v", err)
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("converted bytes = % X, want % X", got, tt.want)
			}
		})
	}
}
