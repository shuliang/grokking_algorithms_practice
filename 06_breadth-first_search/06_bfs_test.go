package main

import (
	"testing"
)

func TestPersonIsSeller(t *testing.T) {
	cases := []struct {
		name string
		want bool
	}{
		{
			name: "you",
			want: false,
		},
		{
			name: "peggy",
			want: false,
		},
		{
			name: "thom",
			want: true,
		},
		{
			name: "thmo",
			want: false,
		},
		{
			name: "",
			want: false,
		},
	}

	for _, c := range cases {
		if got := personIsSeller(c.name); got != c.want {
			t.Errorf("check if '%v' is seller, want %v, got %v ", c.name, c.want, got)
		}
	}
}

func TestBfs(t *testing.T) {
	var graph = map[string][]string{
		"you":    []string{"alice", "bob", "claire"},
		"bob":    []string{"anuj", "peggy"},
		"alice":  []string{"you", "peggy", "alice"},
		"claire": []string{"jonny", "thom", "bob"},
		"anuj":   []string{},
		"peggy":  []string{"you"},
		"thom":   []string{""},
		"jonny":  []string{"peggy", "peggy"},
		"jim":    []string{"thom"},
	}

	cases := []struct {
		start, want string
	}{
		{start: "you", want: "thom"},
		{start: "anuj", want: ""},
		{start: "jonny", want: "thom"},
		{start: "thom", want: ""},
		{start: "jim", want: "thom"},
	}

	for _, c := range cases {
		if got := bfs(graph, c.start); got != c.want {
			t.Errorf("bfs for '%v', want '%v', got '%v'", c.start, c.want, got)
		}
	}
}
