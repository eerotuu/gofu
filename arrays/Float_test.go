package arrays

import (
	"fmt"
	"testing"
)

func TestFloatHas(t *testing.T) {
	data := Float{0.1, 4.3, 2.5}
	if got := data.Has(4.3); !got {
		t.Errorf("Has(3) = %t; want true", got)
	}

	if got := data.Has(6.9); got {
		t.Errorf("Has(6) = %t; want false", got)
	}
}

func TestFloatFilter(t *testing.T) {
	data := Float{0.1, 4.3, 2.5}
	want := Float{0.1, 4.3}
	got := data.Filter(func(i float32) bool {
		return i != 2.5
	})

	for i, v := range got {
		if v != want[i] {
			t.Errorf("Filter() = %f; want %f", got, want)
			break
		}
	}
}

func TestFloatMap(t *testing.T) {
	data := Float{0.1, 4.3, 2.5}
	want := Float{0.2, 8.6, 5.0}
	got := data.Map(func(i float32) float32 {
		return i + i
	})

	for i, v := range got {
		if v != want[i] {
			t.Errorf("Map() = %f; want %f", got, want)
			break
		}
	}
}

func TestFloatReduce(t *testing.T) {
	data := Float{0.1, 4.3, 2.5}
	want := Float{0, 4.3, 5.0}
	got := data.Reduce(func(acc Float, curr float32) Float {
		return append(acc, curr*float32(len(acc)))
	})

	for i, v := range got {
		if v != want[i] {
			t.Errorf("Reduce() = %f; want %f", got, want)
			break
		}
	}
}

func TestFloatStr(t *testing.T) {
	data := Float{1.123, 0.1233, 5123.123}
	want := Str{"1.123", "0.1233", "5123.123"}
	got := data.Str()

	if l := len(got); l != len(want) {
		t.Errorf("Converted array length doesn't match: got = %d, want = 3", l)
	}

	for i, v := range got {
		if v != want[i] {
			t.Errorf("Int()[%d] = %s, want %s", i, v, want[i])
		}
	}
}

func TestFloatInt(t *testing.T) {
	data := Float{1.123, 0.1233, 5123.823}
	want := Int{1, 0, 5123}
	got := data.Int()

	if l := len(got); l != len(want) {
		t.Errorf("Converted array length doesn't match: got = %d, want = 3", l)
	}

	for i, v := range got {
		if v != want[i] {
			t.Errorf("Int()[%d] = %d, want %d", i, v, want[i])
		}
	}
}

func TestFloatDouble(t *testing.T) {
	data := Float{1.1, 0.1233, 5123.823}
	want := Double{1.1, 0.1233, 5123.823}
	got := data.Double()

	x := float32(1.123)
	y := float64(1.123)
	fmt.Printf("%g, %g\n", x, y)

	if l := len(got); l != len(want) {
		t.Errorf("Converted array length doesn't match: got = %d, want = 3", l)
	}

	// Go uses IEEE-754 binary floating-point numbers. Values are not exaclty
	// the same after converting from float32 to float64
}
