package sum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	vals := []int{1, 2, 3, 4, 5}
	vals1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(reflect.TypeOf(vals))
	fmt.Printf("%T\n", vals1)
	fmt.Println(sum(vals...))
}
