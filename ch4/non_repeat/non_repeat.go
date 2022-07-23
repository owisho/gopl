package non_repeat

func RemoveRepeat(s []string) []string {
	count := 0
	pre := ""
	index := 0
	for _, tmp := range s {
		if tmp != pre {
			pre = tmp
			s[index] = tmp
			index++
		} else {
			count++
		}
	}
	return s[:len(s)-count]
}
