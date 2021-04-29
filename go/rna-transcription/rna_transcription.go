package strand

func ToRNA(dna string) string {
	var out string
	for _, d := range dna {
		switch d {
		case 'G':
			out += "C"
		case 'C':
			out += "G"
		case 'T':
			out += "A"
		case 'A':
			out += "U"
		}
	}
	return out
}
