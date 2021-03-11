package arrays

import (
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
