package repeat

func Repeat(r string) (s string) {
	for i := 0; i < 5; i++ {
		s += r
	}
	return s
}
