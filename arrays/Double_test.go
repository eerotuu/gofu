package arrays

import (
	"testing"
)

func TestDoubleHas(t *testing.T) {
	data := Double{0.1, 4.3, 2.5}
	if got := data.Has(4.3); !got {
		t.Errorf("Has(3) = %t; want true", got)
	}

	if got := data.Has(6.9); got {
		t.Errorf("Has(6) = %t; want false", got)
	}
}

func TestDoubleFilter(t *testing.T) {
	data := Double{0.1, 4.3, 2.5}
	want := Double{0.1, 4.3}
	got := data.Filter(func(i float64) bool {
		return i != 2.5
	})

	for i, v := range got {
		if v != want[i] {
			t.Errorf("Filter() = %f; want %f", got, want)
			break
		}
	}
}

func TestDoubleMap(t *testing.T) {
	data := Double{0.1, 4.3, 2.5}
	want := Double{0.2, 8.6, 5.0}
	got := data.Map(func(i float64) float64 {
		return i + i
	})

	for i, v := range got {
		if v != want[i] {
			t.Errorf("Map() = %f; want %f", got, want)
			break
		}
	}
}

func TestDoubleReduce(t *testing.T) {
	data := Double{0.1, 4.3, 2.5}
	want := Double{0, 4.3, 5.0}
	got := data.Reduce(func(acc Double, curr float64) Double {
		return append(acc, curr*float64(len(acc)))
	})

	for i, v := range got {
		if v != want[i] {
			t.Errorf("Reduce() = %f; want %f", got, want)
			break
		}
	}
}
