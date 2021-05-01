package accumulate

func Accumulate(s []string, f func(string) string) []string {
	out := make([]string, len(s))
	for i, ss := range s {
		out[i] = f(ss)
	}
	return out
}
