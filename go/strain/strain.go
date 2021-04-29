package strain

type Ints []int
type Lists [][]int
type Strings []string

func (il Ints) Keep(f func(int) bool) (o Ints) {
	for _, v := range il {
		if f(v) {
			o = append(o, v)
		}
	}
	return
}

func (il Ints) Discard(f func(int) bool) (o Ints) {
	for _, v := range il {
		if !f(v) {
			o = append(o, v)
		}
	}
	return
}

func (il Lists) Keep(f func([]int) bool) (o Lists) {
	for _, v := range il {
		if f(v) {
			o = append(o, v)
		}
	}
	return
}

func (ss Strings) Keep(f func(string) bool) (o Strings) {
	for _, v := range ss {
		if f(v) {
			o = append(o, v)
		}
	}
	return
}
