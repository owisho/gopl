package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/Users/wangyang/Desktop/1.txt")
	if err != nil {
		fmt.Printf("err=%v\n", err)
		return
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}
