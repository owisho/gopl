package max_min

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	vals := []int{1, 2, 3, 4, 5}
	fmt.Println(max(vals...))

}

func TestMin(t *testing.T) {
	vals := []int{1, 2, 3, 4, 5}
	fmt.Println(min(vals...))
}

func TestMultiJoin(t *testing.T) {
	strs := []string{"hello", "xxxx", "world"}
	fmt.Println(multiJoin(",", strs...))
}
