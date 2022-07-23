package rev

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverse1(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseBr(br []byte) {
	//rr := []rune(string(br))
	//for i, j := 0, len(rr)-1; i < j; i, j = i+1, j-1 {
	//	rr[i], rr[j] = rr[j], rr[i]
	//}
	for i, j := 0, len(br)-1; i < j; i, j = i+1, j-1 {
		br[i], br[j] = br[j], br[i]
	}
}
