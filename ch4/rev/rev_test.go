package rev

import (
	"fmt"
	"testing"
)

var a = [...]int{0, 1, 2, 3, 4, 5}

func TestReverse(t *testing.T) {
	fmt.Printf("org=%v\n", a)
	reverse(a[:])
	fmt.Printf("reverse=%v\n", a)
}

func TestReverse1(t *testing.T) {
	fmt.Printf("org=%v\n", a)
	reverse1(&a)
	fmt.Printf("reverse1=%v\n", a)
}
