package arrays

type Float []float32

func (F Float) Has(f float32) bool {
	for _, v := range F {
		if v == f {
			return true
		}
	}

	return false
}

func (F Float) Filter(cond func(f float32) bool) (r Float) {
	for _, v := range F {
		if cond(v) {
			r = append(r, v)
		}
	}

	return r
}

func (F Float) Map(fu func(f float32) float32) (r Float) {
	for _, v := range F {
		r = append(r, fu(v))
	}

	return r
}

func (F Float) Reduce(reducer func(acc Float, curr float32) Float) (r Float) {
	for _, v := range F {
		r = reducer(r, v)
	}

	return r
}

func (F Float) DropLast() Float {
	return F[:len(F)-1]
}
