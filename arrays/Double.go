package arrays

type Double []float64

func (F Double) Has(f float64) bool {
	for _, v := range F {
		if v == f {
			return true
		}
	}

	return false
}

func (F Double) Filter(cond func(f float64) bool) (r Double) {
	for _, v := range F {
		if cond(v) {
			r = append(r, v)
		}
	}

	return r
}

func (F Double) Map(fu func(f float64) float64) (r Double) {
	for _, v := range F {
		r = append(r, fu(v))
	}

	return r
}

func (F Double) Map2Str(fu func(f float64) string) (r Str) {
	for _, v := range F {
		r = append(r, fu(v))
	}

	return r
}

func (F Double) Map2Int(fu func(f float64) int) (r Int) {
	for _, v := range F {
		r = append(r, fu(v))
	}

	return r
}

func (F Double) Map2Float(fu func(f float64) float32) (r Float) {
	for _, v := range F {
		r = append(r, fu(v))
	}

	return r
}

func (F Double) Reduce(reducer func(acc Double, curr float64) Double) (r Double) {
	for _, v := range F {
		r = reducer(r, v)
	}

	return r
}

func (F Double) DropLast() Double {
	return F[:len(F)-1]
}
