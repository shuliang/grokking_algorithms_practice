package main

import (
	"reflect"
	"testing"
)

func TestDijkstra(t *testing.T) {
	g1 := graph{
		"start": map[string]int{"a": 6, "b": 2},
		"a":     map[string]int{"fin": 1},
		"b":     map[string]int{"a": 3, "fin": 5},
		"fin":   map[string]int{},
	}

	g2 := graph{
		"start": map[string]int{"a": 6, "b": 2, "d": 1},
		"a":     map[string]int{"c": 2},
		"b":     map[string]int{"a": 3, "fin": 5},
		"c":     map[string]int{"fin": 2},
		"d":     map[string]int{"b": 4, "a": 3},
		"fin":   map[string]int{},
	}

	// graph contains circle
	g3 := graph{
		"start": map[string]int{"a": 6, "b": 2},
		"a":     map[string]int{"b": 2},
		"b":     map[string]int{"a": 3, "fin": 15},
		"fin":   map[string]int{},
	}

	// graph without "fin"
	g4 := graph{
		"start": map[string]int{"a": 6, "b": 2},
		"a":     map[string]int{"b": 2},
		"b":     map[string]int{},
	}

	// graph with negative cost
	g5 := graph{
		"start": map[string]int{"a": 6, "b": 2, "d": 5},
		"a":     map[string]int{"c": -2, "fin": 1},
		"b":     map[string]int{"a": -3, "fin": 5},
		"c":     map[string]int{"fin": 2},
		"d":     map[string]int{"b": -4, "a": 3},
		"fin":   map[string]int{},
	}

	cases := []struct {
		g     graph
		start string
		want  []string
	}{
		{
			g:     g1,
			start: "start",
			want:  []string{"start", "b", "a", "fin"},
		},
		{
			g:     g1,
			start: "b",
			want:  []string{"b", "a", "fin"},
		},
		{
			g:     g1,
			start: "a",
			want:  []string{"a", "fin"},
		},
		{
			g:     g1,
			start: "wrongStart",
			want:  nil,
		},
		{
			g:     g2,
			start: "start",
			want:  []string{"start", "b", "fin"},
		},
		{
			g:     g2,
			start: "d",
			want:  []string{"d", "a", "c", "fin"},
		},
		{
			g:     g3,
			start: "start",
			want:  []string{"start", "b", "fin"},
		},
		{
			g:     g4,
			start: "start",
			want:  nil,
		},
		{
			g:     g5,
			start: "start",
			want:  []string{"start", "d", "b", "a", "c", "fin"},
		},
		{
			g:     g5,
			start: "d",
			want:  []string{"d", "b", "a", "c", "fin"},
		},
	}

	for _, c := range cases {
		if got := dijkstra(c.g, c.start); !reflect.DeepEqual(c.want, got) {
			t.Errorf("Find shortest path of graph '%v', start: '%s', want: '%v', got: '%v'", c.g, c.start, c.want, got)
		}
	}
}
