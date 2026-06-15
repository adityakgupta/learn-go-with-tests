package sum

func Sum(a []int) (r int) {
	for _, num := range a {
		r += num
	}
	return
}

func SumAll(numbers ...[]int) (sum []int) {
	for _, n := range numbers {
		sum = append(sum, Sum(n))
	}

	return
}

func SumAllTails(numbers ...[]int) (sum []int) {
	for _, n := range numbers {
		if len(n) == 0 {
			sum = append(sum, 0)
		} else {
			t := n[1:]
			sum = append(sum, Sum(t))
		}
	}
	return
}
