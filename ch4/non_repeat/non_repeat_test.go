package non_repeat

import (
	"fmt"
	"testing"
)

func TestRemoveRepeat(t *testing.T) {
	s := []string{"1", "1", "2", "2", "3"}
	fmt.Printf("before=%v\n", s)
	s = RemoveRepeat(s)
	fmt.Printf("after=%v\n", s)
}
