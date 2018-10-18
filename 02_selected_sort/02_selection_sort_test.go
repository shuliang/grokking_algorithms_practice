package main

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	cases := []struct {
		origin []int
		want   []int
	}{
		{
			origin: []int{1, 8, 5, 9, 7, 6, 3},
			want:   []int{1, 3, 5, 6, 7, 8, 9},
		},
		{
			origin: []int{8, 4, 6, 2},
			want:   []int{2, 4, 6, 8},
		},
		{
			origin: []int{5, 3, 6, 2, 0},
			want:   []int{0, 2, 3, 5, 6},
		},
	}

	for _, c := range cases {
		got := selectionSort(c.origin)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("sort %v, want %v, got %v", c.origin, c.want, got)
		}
	}
}
