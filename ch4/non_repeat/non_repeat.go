package non_repeat

import "unicode"

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

func RemoveDupSpace(br []byte) []byte {
	if len(br) == 0 {
		return br
	}
	rs := []rune(string(br))
	i := 0
	isSpace := false
	for _, r := range rs {
		if isSpace && unicode.IsSpace(r) {
			continue
		}
		isSpace = false
		if unicode.IsSpace(r) {
			isSpace = true
		}
		rs[i] = r
		i++
	}
	return []byte(string(rs[:i]))
}
