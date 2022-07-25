package tmp

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m := make(map[string]int)
	i, ok := m["test"]
	fmt.Printf("i=%d,ok=%t\n", i, ok)
	m["test"] = 0
	i, ok = m["test"]
	fmt.Printf("i=%d,ok=%t\n", i, ok)
}
