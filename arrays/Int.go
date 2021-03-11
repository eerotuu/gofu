package arrays

import "fmt"

type Int []int

func (I Int) Has(i int) bool {
	for _, v := range I {
		if v == i {
			return true
		}
	}

	return false
}

func (I Int) Filter(cond func(i int) bool) (r Int) {
	for _, v := range I {
		if cond(v) {
			r = append(r, v)
		}
	}

	return r
}

func (I Int) Map(fu func(i int) int) (r Int) {
	for _, v := range I {
		r = append(r, fu(v))
	}

	return r
}

func (I Int) Map2Str(fu func(i int) string) (r Str) {
	for _, v := range I {
		r = append(r, fu(v))
	}

	return r
}

func (I Int) Map2Float(fu func(i int) float32) (r Float) {
	for _, v := range I {
		r = append(r, fu(v))
	}

	return r
}

func (I Int) Map2Double(fu func(i int) float64) (r Double) {
	for _, v := range I {
		r = append(r, fu(v))
	}

	return r
}

func (I Int) Reduce(reducer func(acc Int, curr int) Int) (r Int) {
	for _, v := range I {
		r = reducer(r, v)
	}

	return r
}

func (I Int) DropLast() Int {
	return I[:len(I)-1]
}

func (I Int) Str() (r Str) {
	r = Str{}
	for _, v := range I {
		r = append(r, fmt.Sprint(v))
	}

	return r
}

func (I Int) Float() (r Float) {
	r = Float{}
	for _, v := range I {
		r = append(r, float32(v))
	}

	return r
}

func (I Int) Double() (r Double) {
	r = Double{}
	for _, v := range I {
		r = append(r, float64(v))
	}

	return r
}
