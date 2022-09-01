package ch6

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	c := Counter{n: 1}
	fmt.Println(c.N())
	c.n = 100
	fmt.Println(c.N())
}
