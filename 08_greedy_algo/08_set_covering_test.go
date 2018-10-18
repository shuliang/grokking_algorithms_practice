package main

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	cases := []struct {
		a, b, want []string
	}{
		{
			a:    []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"},
			b:    []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"},
			want: []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"},
		},
		{
			a:    []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"},
			b:    []string{"id", "nv", "ut", "ca", "az", "mt", "wa", "or"},
			want: []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"},
		},
		{
			a:    []string{"wa", "or", "id", "nv"},
			b:    []string{"mt", "wa", "nv", "ut", "ca", "az"},
			want: []string{"wa", "nv"},
		},
		{
			a:    []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"},
			b:    []string{"amt", "awa", "aor", "Id", "Nv", "Ut", "Ca", "aZ"},
			want: []string{},
		},
		{
			a:    []string{},
			b:    []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"},
			want: []string{},
		},
		{
			a:    []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"},
			b:    []string{},
			want: []string{},
		},
	}
	for _, c := range cases {
		if got := intersect(c.a, c.b); !reflect.DeepEqual(c.want, got) {
			t.Errorf("intersect of %v and %v, want %v, got %v", c.a, c.b, c.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	cases := []struct {
		a, b, want []string
	}{
		{
			a:    []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"},
			b:    []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"},
			want: []string{},
		},
		{
			a:    []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"},
			b:    []string{"id", "nv", "ut", "ca", "az", "mt", "wa", "or"},
			want: []string{},
		},
		{
			a:    []string{"wa", "or", "id", "nv"},
			b:    []string{"mt", "wa", "nv", "ut", "ca", "az"},
			want: []string{"or", "id"},
		},
		{
			a:    []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"},
			b:    []string{"amt", "awa", "aor", "Id", "Nv", "Ut", "Ca", "aZ"},
			want: []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"},
		},
		{
			a:    []string{},
			b:    []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"},
			want: []string{},
		},
		{
			a:    []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"},
			b:    []string{},
			want: []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"},
		},
	}
	for _, c := range cases {
		if got := subtract(c.a, c.b); !reflect.DeepEqual(c.want, got) {
			t.Errorf("%v subtract %v, want %v, got %v", c.a, c.b, c.want, got)
		}
	}
}
