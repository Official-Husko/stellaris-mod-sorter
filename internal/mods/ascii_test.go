package mods

import "testing"

func TestToASCII(t *testing.T) {
	cases := []struct {
		in  string
		out string
	}{
		{"abc", "abc"},
		{"héllo世界", "hllo"},
		{"123!@#", "123!@#"},
		{"", ""},
		{"こんにちは", ""},
		{"Go语言123", "Go123"},
	}
	for _, c := range cases {
		if got := toASCII(c.in); got != c.out {
			t.Errorf("toASCII(%q) = %q, want %q", c.in, got, c.out)
		}
	}
}
