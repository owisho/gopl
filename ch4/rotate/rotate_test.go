package rotate

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("org=%v\n", s)
	rotate(s, 3)
	fmt.Printf("rotate=%v\n", s)
}
