package main

import (
	"fmt"
	"go/types"
)

func main() {
	s := "国"
	fmt.Println(len(s))
	a1 := s[1:]
	fmt.Println(a1)
	a2 := make([]byte, 0)
	a2 = append(a2, s...)
	fmt.Println(a2)
	a3 := [3]byte{1, 2, 3}
	a4 := []byte{1, 2, 3}
	fmt.Printf("a1 type = %s\n", printType(a1))
	fmt.Printf("a2 type = %s\n", printType(a2))
	fmt.Printf("a3 type = %s\n", printType(a3))
	fmt.Printf("a4 type = %s\n", printType(a4))

	b := []byte(s)
	fmt.Println(b)
	b[2] = 188
	s1 := string(b)
	fmt.Println(s, s1)

	symbol := [...]string{2: "ssss", 3: "xxxx"}
	fmt.Println(len(symbol))

	s2 := "asdfg"
	s2 = s2[1:2]
	fmt.Println(s2)

}

func printType(a interface{}) string {
	switch a.(type) {
	case []byte:
		return "byte slice"
	case types.Slice:
		return "slice"
	case string:
		return "string"
	default:
		return "unknown"
	}
}

func t() {
	var cp complex128
	cp = complex(float64(1.0), float64(2.0))
	fmt.Println(cp)
}
