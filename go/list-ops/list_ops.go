package listops

type IntList []int
type binFunc func(x, y int) int
type predFunc func(x int) bool
type unaryFunc func(x int) int

func (l IntList) Length() int {
	length := 0
	for range l {
		length++
	}
	return length
}

func (l IntList) Reverse() IntList {
	last := l.Length() - 1
	var tmp int
	for i := 0; i < last; i++ {
		tmp = l[last]
		l[last] = l[i]
		l[i] = tmp
		last--
	}
	return l
}

func (a IntList) Append(b IntList) IntList {
	a_len := a.Length()
	out := make(IntList, a_len+b.Length())
	for i, v := range a {
		out[i] = v
	}
	for i, v := range b {
		out[a_len+i] = v
	}
	return out
}

func (l IntList) Foldr(f binFunc, i int) int {
	for j := l.Length() - 1; j >= 0; j-- {
		i = f(l[j], i)
	}
	return i
}

func (l IntList) Foldl(f binFunc, i int) int {
	for _, v := range l {
		i = f(i, v)
	}
	return i
}

func (in IntList) Concat(ls []IntList) IntList {
	for _, l := range ls {
		in = in.Append(l)
	}
	return in
}

func (l IntList) Filter(f predFunc) IntList {
	list := make(IntList, 0, l.Length())
	i := 0
	for _, v := range l {
		if f(v) {
			list = list[:i+1]
			list[i] = v
			i++
		}
	}
	return list
}

func (l IntList) Map(f unaryFunc) IntList {
	for i, v := range l {
		l[i] = f(v)
	}
	return l
}
