package rotate

//0, 1, 2, 3, 4, 5
//2 3 4 5 0 1

func rotate(s []int, i int) {
	tmp := make([]int, i)
	for j := 0; j < len(s); j++ {
		if j < i {
			tmp[j] = s[j]
		}
		if j <= len(s)-1-i {
			s[j] = s[j+i]
		} else {
			s[j] = tmp[j+i-len(s)]
		}
	}
}
