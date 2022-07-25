package treesort

import (
	"fmt"
	"testing"
)

func TestTreesort(t *testing.T) {
	s := []int{2, 9, 1, 4, 5, 3, 8}
	fmt.Println(s)
	Sort(s)
	fmt.Println(s)
}
