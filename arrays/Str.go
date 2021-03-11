package arrays

import (
	"strconv"
)

// The Concat method is used to merge two or more strings.
func Concat(s ...string) (r string) {
	for _, v := range s {
		r += v
	}

	return r
}

type Str []string

// The Has() method checks if Str contains provided value
func (S Str) Has(s string) bool {
	for _, v := range S {
		if v == s {
			return true
		}
	}
	return false
}

// The Map() method creates a new Str populated with the results of calling a provided
// function on every element in the calling array.
func (S Str) Map(mapper func(s string) string) (r Str) {
	for _, v := range S {
		r = append(r, mapper(v))
	}

	return r
}

func (S Str) Map2Int(mapper func(s string) int) (r Int) {
	for _, v := range S {
		r = append(r, mapper(v))
	}

	return r
}

func (S Str) Map2Float(mapper func(s string) float32) (r Float) {
	for _, v := range S {
		r = append(r, mapper(v))
	}

	return r
}

func (S Str) Map2Double(mapper func(s string) float64) (r Double) {
	for _, v := range S {
		r = append(r, mapper(v))
	}

	return r
}

// The Filter() method creates a new Str with all elemetns taht pass the test implemented
// by the provided function.
func (S Str) Filter(cond func(s string) bool) (r Str) {
	for _, v := range S {
		if cond(v) {
			r = append(r, v)
		}
	}

	return r
}

// The Reduce() method executes a reducer function (that you provide) on element
// of the array, resulting in reduced Str output.
func (S Str) Reduce(reducer func(acc Str, curr string) Str) (r Str) {
	for _, v := range S {
		r = reducer(r, v)
	}

	return r
}

// The Join() method inserts separator between each element and concatenates all the elements,
// resulting in single output string.
func (S Str) Join(separator string) (r string) {
	r = S[0]
	for _, v := range S[1:] {
		r = Concat(r, separator, v)
	}
	return r
}

func (strings Str) DropLast() Str {
	return strings[:len(strings)-1]
}

func (I Str) Int() (r Int, err error) {
	r = Int{}
	for _, v := range I {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		r = append(r, int(i))
	}

	return r, nil
}

func (I Str) Float() (r Float, err error) {
	r = Float{}
	for _, v := range I {
		i, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return nil, err
		}
		r = append(r, float32(i))
	}

	return r, nil
}

func (I Str) Double() (r Double, err error) {
	r = Double{}
	for _, v := range I {
		i, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}
		r = append(r, i)
	}

	return r, nil
}
