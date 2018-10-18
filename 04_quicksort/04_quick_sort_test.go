package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	cases := []struct {
		list      []int
		wantSum   int
		wantCount int
	}{
		{
			list:      []int{1, 29, 20, 38, 14, 5, 72, 3, 3, 2},
			wantSum:   187,
			wantCount: 10,
		},
		{
			list:      []int{-10, 9, 0, 7},
			wantSum:   6,
			wantCount: 4,
		},
	}

	t.Run("loop sum", func(t *testing.T) {
		for _, c := range cases {
			if got := loopSum(c.list); got != c.wantSum {
				t.Errorf("loop sum %v, wang %d, got %d", c.list, c.wantSum, got)
			}
		}
	})
	t.Run("recursive sum", func(t *testing.T) {
		for _, c := range cases {
			if got := recursiveSum(c.list); got != c.wantSum {
				t.Errorf("recursive sum %v, wang %d, got %d", c.list, c.wantSum, got)
			}
		}
	})
	t.Run("recursive count", func(t *testing.T) {
		for _, c := range cases {
			if got := count(c.list); got != c.wantCount {
				t.Errorf("recursive count %v, wang %d, got %d", c.list, c.wantCount, got)
			}
		}
	})
}

func TestQuickSort(t *testing.T) {
	cases := []struct {
		list []int
		want []int
	}{
		{
			list: []int{1, 29, 20, 38, 14, 5, 72, 3, 3, 2},
			want: []int{1, 2, 3, 3, 5, 14, 20, 29, 38, 72},
		},
		{
			list: []int{-10, 9, 0, 7},
			want: []int{-10, 0, 7, 9},
		},
	}

	t.Run("quick sort pivot middle", func(t *testing.T) {
		for _, c := range cases {
			if got := quickSort(c.list); !reflect.DeepEqual(got, c.want) {
				t.Errorf("quick sort - pivot middle %v, wang %d, got %d", c.list, c.want, got)
			}
		}
	})

	t.Run("quick sort pivot first", func(t *testing.T) {
		for _, c := range cases {
			if got := qSortPivotFirst(c.list); !reflect.DeepEqual(got, c.want) {
				t.Errorf("quick sort - pivot first %v, wang %d, got %d", c.list, c.want, got)
			}
		}
	})
}

func genRandomList(count int) []int {
	list := []int{}
	for i := 0; i < count; i++ {
		list = append(list, rand.Int())
	}
	return list
}

func BenchmarkQuickSortPivotMiddle(b *testing.B) {
	cases := [][]int{
		[]int{1, 29, 20, 38, 14, 5, 72, 3, 3, 2, -10, 67, 23},
		[]int{-10, 9, 0, 7},
		// genRandomList(1000000),
	}
	for _, list := range cases {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			quickSort(list)
		}
	}
}

func BenchmarkQuickSortPivotFirst(b *testing.B) {
	cases := [][]int{
		[]int{1, 29, 20, 38, 14, 5, 72, 3, 3, 2, -10, 67, 23},
		[]int{-10, 9, 0, 7},
		// genRandomList(1000000),
	}
	for _, list := range cases {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			qSortPivotFirst(list)
		}
	}
}
