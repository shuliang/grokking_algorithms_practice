package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		list   []int
		target int
		want   int
	}{
		{
			list:   []int{0, 1, 2, 3, 4, 5, 6, 7},
			target: 0,
			want:   0,
		},
		{
			list:   []int{0, 1, 2, 3, 4, 5, 6, 7},
			target: 7,
			want:   7,
		},
		{
			list:   []int{0, 1, 2, 3, 4, 5, 6, 7},
			target: 3,
			want:   3,
		},
		{
			list:   []int{0, 1, 2, 3, 4, 5, 6, 7},
			target: 4,
			want:   4,
		},
		{
			list:   []int{0, 1, 2, 3, 4, 5, 6, 7},
			target: 5,
			want:   5,
		},
		{
			list:   []int{0, 1, 2, 3, 4, 5, 6, 7},
			target: 6,
			want:   6,
		},
		{
			list:   []int{1, 3, 5, 7, 9},
			target: 1,
			want:   0,
		},
		{
			list:   []int{1, 3, 5, 7, 9},
			target: 5,
			want:   2,
		},
		{
			list:   []int{1, 3, 5, 7, 9},
			target: 7,
			want:   3,
		},
		{
			list:   []int{1, 3, 5, 7, 9},
			target: 9,
			want:   4,
		},
		{
			list:   []int{1, 3, 5, 7, 9},
			target: 10,
			want:   -1,
		},
		{
			list:   []int{2, 4, 6, 8},
			target: 4,
			want:   1,
		},
		{
			list:   []int{2, 4, 6, 8},
			target: 6,
			want:   2,
		},

		{
			list:   []int{2, 4, 6, 8},
			target: 8,
			want:   3,
		},
	}

	for _, c := range cases {
		if got := binarySearch(c.list, c.target); got != c.want {
			t.Errorf("find %d in %v, want %d, got %d", c.target, c.list, c.want, got)
		}
	}
}
