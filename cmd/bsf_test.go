package cmd

import (
	"testing"
)

func Test_toEscape(t *testing.T) {
	tests := []struct {
		input rune
		want  string
	}{
		{'A', "\\u0041"},
		{' ', "\\u0020"},
		{0xFFFF, "\\uffff"},
		{'中', "\\u4e2d"},
	}
	for _, tt := range tests {
		if got := toEscape(tt.input); got != tt.want {
			t.Errorf("toEscape(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func Test_fromEscaped(t *testing.T) {
	tests := []struct {
		name  string
		input []rune
		want  string
	}{
		{
			"valid uf240",
			[]rune{'u', 'f', '2', '4', '0'},
			string(rune(0xf240)),
		},
		{
			"valid uppercase hex",
			[]rune{'u', 'F', '2', '4', '0'},
			string(rune(0xf240)),
		},
		{
			"too short",
			[]rune{'u', 'f', '2', '4'},
			"Invalid hexadecimal character",
		},
		{
			"too long",
			[]rune{'u', 'f', '2', '4', '0', 'a'},
			"Invalid hexadecimal character",
		},
		{
			"invalid hex char",
			[]rune{'u', 'f', '2', '4', 'g'},
			"Invalid format",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fromEscaped(tt.input)
			if got != tt.want {
				t.Errorf("fromEscaped(%v) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func Test_format(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"plain string", "abc", "abc"},
		// 6-char literal ä; strconv.Unquote decodes the unicode escape to ä
		{"unicode escape sequence", "\\u00e4", "ä"},
		// \xgg is invalid because g is not a hex digit
		{"invalid escape", "\\xgg", "Failed to unquote string"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := format(tt.input); got != tt.want {
				t.Errorf("format(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
