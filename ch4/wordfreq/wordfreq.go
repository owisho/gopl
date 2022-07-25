package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		counts[word]++
		if word == "end" {
			break
		}
	}
	fmt.Printf("word\tcount\n")
	for word, count := range counts {
		fmt.Printf("%s\t%d\n", word, count)
	}
}
