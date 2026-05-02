package cmd

import (
	"bytes"
	"testing"
)

func Test_get_lf_index(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  []int
	}{
		{"empty", nil, nil},
		{"no newline", []byte("abc"), nil},
		{"single bare LF", []byte("a\nb"), []int{1}},
		{"CRLF ignored", []byte("a\r\nb"), nil},
		{"bare LF before CRLF", []byte("a\nb\r\nc"), []int{1}},
		{"CRLF before bare LF", []byte("a\r\nb\nc"), []int{4}},
		{"multiple bare LFs", []byte("a\nb\nc"), []int{1, 3}},
		{"leading LF", []byte("\nabc"), []int{0}},
		{"trailing LF", []byte("abc\n"), []int{3}},
		{"UTF-8 multibyte before LF", []byte("\xe4\xb8\xad\n"), []int{3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := get_lf_index(tt.input)
			if len(got) != len(tt.want) {
				t.Fatalf("get_lf_index: got %v, want %v", got, tt.want)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("index %d: got %d, want %d", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func Test_get_crlf_index(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  []int
	}{
		{"empty", nil, nil},
		{"no newline", []byte("abc"), nil},
		{"bare LF ignored", []byte("a\nb"), nil},
		{"single CRLF", []byte("a\r\nb"), []int{1}},
		{"lone CR at EOF", []byte("a\r"), nil},
		{"multiple CRLFs", []byte("a\r\nb\r\nc"), []int{1, 4}},
		{"UTF-8 before CRLF", []byte("\xe4\xb8\xad\r\n"), []int{3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := get_crlf_index(tt.input)
			if len(got) != len(tt.want) {
				t.Fatalf("get_crlf_index: got %v, want %v", got, tt.want)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("index %d: got %d, want %d", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func Test_lf_to_crlf(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  []byte
	}{
		{"empty", []byte{}, []byte{}},
		{"no LF", []byte("abc"), []byte("abc")},
		{"single LF", []byte("a\nb"), []byte("a\r\nb")},
		{"multiple LFs", []byte("a\nb\nc"), []byte("a\r\nb\r\nc")},
		{"leading LF", []byte("\nabc"), []byte("\r\nabc")},
		{"trailing LF", []byte("abc\n"), []byte("abc\r\n")},
		{"CRLF left unchanged", []byte("a\r\nb"), []byte("a\r\nb")},
		{"mixed CRLF and bare LF", []byte("a\r\nb\nc"), []byte("a\r\nb\r\nc")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lf_to_crlf(tt.input)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("lf_to_crlf(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func Test_crlf_to_lf(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  []byte
	}{
		{"empty", []byte{}, []byte{}},
		{"no CRLF", []byte("abc"), []byte("abc")},
		{"single CRLF", []byte("a\r\nb"), []byte("a\nb")},
		{"multiple CRLFs", []byte("a\r\nb\r\nc"), []byte("a\nb\nc")},
		{"leading CRLF", []byte("\r\nabc"), []byte("\nabc")},
		{"trailing CRLF", []byte("abc\r\n"), []byte("abc\n")},
		{"bare LF left unchanged", []byte("a\nb"), []byte("a\nb")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := crlf_to_lf(tt.input)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("crlf_to_lf(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func Test_roundtrip_lf_crlf(t *testing.T) {
	inputs := [][]byte{
		[]byte("hello\nworld\n"),
		[]byte("line1\nline2\nline3"),
		[]byte("\xe4\xb8\xad\n\xe6\x96\x87"), // 中\n文
	}
	for _, in := range inputs {
		converted := lf_to_crlf(in)
		back := crlf_to_lf(converted)
		if !bytes.Equal(back, in) {
			t.Errorf("round-trip failed for %q: got %q", in, back)
		}
	}
}
