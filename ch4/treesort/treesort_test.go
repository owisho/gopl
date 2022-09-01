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

func TestPrintTree(t *testing.T) {
	tr := &tree{}
	tr = add(tr, 4)
	tr = add(tr, 2)
	tr = add(tr, 5)
	tr = add(tr, 3)
	tr = add(tr, 7)
	fmt.Print(tr)
}
