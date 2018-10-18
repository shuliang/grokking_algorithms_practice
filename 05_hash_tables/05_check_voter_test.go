package main

import (
	"testing"
)

func TestCheck(t *testing.T) {
	n1, n2 := "tom", "mike"
	cases := []struct {
		tName string
		names []string
		want  []string
	}{
		{
			tName: "case1",
			names: []string{n1, n2, n2},
			want:  []string{approve, approve, reject},
		},
		{
			tName: "case2",
			names: []string{n1, n1, n2, n2},
			want:  []string{approve, reject, approve, reject},
		},
		{
			tName: "case3",
			names: []string{n1, n2, n1, n2},
			want:  []string{approve, approve, reject, reject},
		},
		{
			tName: "case4",
			names: []string{n1, n2, n2, n2, n1, n1},
			want:  []string{approve, approve, reject, reject, reject, reject, reject},
		},
	}

	for _, c := range cases {
		t.Run(c.tName, func(t *testing.T) {
			// must reset voted map here
			voted = map[string]bool{}

			for i, name := range c.names {
				if got := check(name); got != c.want[i] {
					t.Errorf("check '%s' in '%v', want '%s', got '%s'", name, c.names, c.want[i], got)
				}
			}
		})
	}
}
