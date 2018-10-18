package main

import (
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	cases := []struct {
		a, b string
		want string
	}{
		{
			a:    "fish",
			b:    "fosh",
			want: "sh",
		},
		{
			a:    "fioo",
			b:    "fosh",
			want: "f",
		},
		{
			a:    "dish",
			b:    "fo000000sh",
			want: "sh",
		},
		{
			a:    "vista",
			b:    "fosh",
			want: "s",
		},
		{
			a:    "fish",
			b:    "",
			want: "",
		},
		{
			a:    "",
			b:    "fosh",
			want: "",
		},
		{
			a:    "Hello世界",
			b:    "hello界世界",
			want: "ello世界",
		},
		{
			a:    "Hello世79000000000000000界",
			b:    "hello界世界",
			want: "ello世",
		},
		{
			a:    "Hello世7界",
			b:    "hello界世界",
			want: "llo世界",
		},
		{
			a:    "Hello世界世界",
			b:    "hello世界界世",
			want: "ello世界世",
		},
		{
			a:    "Hello世89界7世界",
			b:    "hello世89界7界世",
			want: "ello世89界7世",
		},
		{
			a:    "Hello123世界世界",
			b:    "hello123界世界世",
			want: "ello123世界世",
		},
		{
			a:    "世90界8世界",
			b:    "世90界8界世",
			want: "世90界8世",
		},
	}

	for _, c := range cases {
		if got := longestCommonSubsequence(c.a, c.b); got != c.want {
			t.Errorf("longest subsequence of '%v' and '%v', want:'%v', got:'%v'", c.a, c.b, c.want, got)
		}
	}
}

func TestLongestCommonSubstring(t *testing.T) {
	cases := []struct {
		a, b string
		want string
	}{
		{
			a:    "fish",
			b:    "fosh",
			want: "sh",
		},
		{
			a:    "fioo",
			b:    "fosh",
			want: "f",
		},
		{
			a:    "dish",
			b:    "fo000000sh",
			want: "sh",
		},
		{
			a:    "vista",
			b:    "fosh",
			want: "s",
		},
		{
			a:    "fish",
			b:    "",
			want: "",
		},
		{
			a:    "",
			b:    "fosh",
			want: "",
		},
		{
			a:    "Hello世界",
			b:    "hello界",
			want: "ello",
		},
		{
			a:    "Hello世79000000000000000界",
			b:    "hello世界",
			want: "ello世",
		},
		{
			a:    "Hello世界世界",
			b:    "hello世界界世",
			want: "ello世界",
		},
		{
			a:    "Hello世89界7世界",
			b:    "hello世89界7界世",
			want: "ello世89界7",
		},
		{
			a:    "Hello123世界世界",
			b:    "hello123界世界世",
			want: "ello123",
		},
		{
			a:    "世90界8世界",
			b:    "世90界8界世",
			want: "世90界8",
		},
	}

	for _, c := range cases {
		if got := longestCommonSubstring(c.a, c.b); got != c.want {
			t.Errorf("longest substring of '%v' and '%v', want:'%v', got:'%v'", c.a, c.b, c.want, got)
		}
	}
}
