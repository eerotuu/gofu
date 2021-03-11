package arrays

import (
	"testing"
)

func TestIntHas(t *testing.T) {
	data := Int{1, 3, 5}
	if got := data.Has(3); !got {
		t.Errorf("Has(3) = %t; want true", got)
	}

	if got := data.Has(6); got {
		t.Errorf("Has(6) = %t; want false", got)
	}
}

func TestIntFilter(t *testing.T) {
	data := Int{1, 3, 5}
	want := Int{3, 5}
	got := data.Filter(func(i int) bool {
		return i != 1
	})

	for i, v := range got {
		if v != want[i] {
			t.Errorf("Filter() = %d; want false", got)
			break
		}
	}

}

func TestIntMap(t *testing.T) {
	data := Int{1, 3, 5}
	want := Int{1, 9, 25}
	got := data.Map(func(i int) int {
		return i * i
	})

	for i, v := range got {
		if v != want[i] {
			t.Errorf("Map() = %d; want [1 9 25]", got)
			break
		}
	}
}

func TestIntReduce(t *testing.T) {
	data := Int{1, 3, 5}
	want := Int{0, 3, 10}
	got := data.Reduce(func(acc Int, curr int) Int {
		return append(acc, curr*len(acc))
	})

	for i, v := range got {
		if v != want[i] {
			t.Errorf("Reduce() = %d; want %d", got, want)
			break
		}
	}
}

func TestIntStr(t *testing.T) {
	data := Int{1, 3, 5}
	want := Str{"1", "3", "5"}
	got := data.Str()

	if l := len(got); l != len(want) {
		t.Errorf("Converted array length doesn't match: got = %d, want = 3", l)
	}

	for i, v := range got {
		if v != want[i] {
			t.Errorf("Str() = %s, want %s", got, want)
			break
		}
	}
}

func TestIntFloat(t *testing.T) {
	data := Int{1, 3, 5}
	want := Float{1, 3, 5}
	got := data.Float()

	if l := len(got); l != len(want) {
		t.Errorf("Converted array length doesn't match: got = %d, want = 3", l)
	}

	for i, v := range got {
		if v != want[i] {
			t.Errorf("Str() = %f, want %f", got, want)
			break
		}
	}
}

func TestIntDouble(t *testing.T) {
	data := Int{1, 3, 5}
	want := Double{1, 3, 5}
	got := data.Double()

	if l := len(got); l != len(want) {
		t.Errorf("Converted array length doesn't match: got = %d, want = 3", l)
	}

	for i, v := range got {
		if v != want[i] {
			t.Errorf("Str() = %f, want %f", got, want)
			break
		}
	}
}
