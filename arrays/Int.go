package arrays

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

func (I Int) Reduce(reducer func(acc Int, curr int) Int) (r Int) {
	for _, v := range I {
		r = reducer(r, v)
	}

	return r
}

func (I Int) DropLast() Int {
	return I[:len(I)-1]
}

func (I Int) Get() []int {
	return []int(I)
}
