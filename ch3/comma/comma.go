package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bufferComma("1231231231"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma1(s string) string {
	n := len(s)
	ret := ""
	if n <= 3 {
		return s
	}
	i := n - 3
	for ; i > 0; i = i - 3 {
		pre := "," + s[i:i+3]
		ret = pre + ret
	}
	ret = s[0:i+3] + ret
	return ret
}

func bufferComma(s string) string {
	n := len(s)
	var bf bytes.Buffer
	count := 0
	for i := n - 1; i >= 0; i-- {
		count++
		bf.WriteByte(s[i])
		if count == 3 && i != 0 {
			bf.WriteByte(',')
			count = 0
		}
	}
	br := bf.Bytes()
	var bf1 bytes.Buffer
	for i := len(br) - 1; i >= 0; i-- {
		bf1.WriteByte(br[i])
	}
	return bf1.String()

}
