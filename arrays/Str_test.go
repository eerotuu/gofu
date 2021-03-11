package arrays

import (
	"strings"
	"testing"
)

func TestStrHas(t *testing.T) {
	data := Str{"apple", "banana", "orange"}
	if got := data.Has("apple"); !got {
		t.Errorf("Has(3) = %t; want true", got)
	}

	if got := data.Has("mango"); got {
		t.Errorf("Has(mango) = %t; want false", got)
	}
}

func TestStrFilter(t *testing.T) {
	data := Str{"banana", "apple", "orange", "orange"}
	want := Str{"banana", "apple"}
	got := data.Filter(func(i string) bool {
		return i != "orange"
	})

	if l := len(got); l != len(want) {
		t.Errorf("len(Filter()) = %d; want 2", l)
	}

	for i, v := range got {
		if v != want[i] {
			t.Errorf("Filter() = %s; want false", got)
			break
		}
	}
}

func TestStrMap(t *testing.T) {
	data := Str{"apple", "banana", "orange"}
	want := Str{"Apple", "Banana", "Orange"}
	got := data.Map(func(s string) string {
		return strings.Title(s)
	})

	if l := len(got); l != len(want) {
		t.Errorf("len(Filter()) = %d; want 3", l)
	}

	for i, v := range got {
		if v != want[i] {
			t.Errorf("Map() = %s; want %s", got, want)
			break
		}
	}
}

func TestStrReduce(t *testing.T) {
	data := Str{"apple", "banana", "orange", "banana"}
	want := Str{"apple", "banana", "orange"}
	got := data.Reduce(func(acc Str, curr string) Str {
		if !acc.Has(curr) {
			return append(acc, curr)
		}

		return acc
	})

	if l := len(got); l != len(want) {
		t.Errorf("len(Filter()) = %d; want 3", l)
	}

	for i, v := range got {
		if v != want[i] {
			t.Errorf("Reduce() = %s; want %s", got, want)
			break
		}
	}
}

func TestStrJoin(t *testing.T) {
	data := Str{"Hello ", "World!"}
	want := "Hello World!"

	if got := data.Join(); got != want {
		t.Errorf("Join() = %s; want %s", got, want)
	}
}
