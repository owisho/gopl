package tmp

import (
	"fmt"
	"testing"
	"time"
)

func TestDefer(t *testing.T) {
	addNum(1, 2)
}
func addNum(a, b int) int {
	t1 := time.Now()
	fmt.Printf("begin time %p\n", &t1)
	fmt.Printf("begin time %v\n", t1)
	//defer func(start time.Time) {
	//	fmt.Printf("param start %p\n", &start)
	//	fmt.Printf("param start %v\n", start)
	//}(t1)
	defer print(t1)
	time.Sleep(3 * time.Second)
	t1 = time.Now()
	fmt.Printf("end sleep %p\n", &t1)
	fmt.Printf("end sleep %v\n", &t1)
	return a + b
}
func print(start time.Time) {
	fmt.Printf("param start %p\n", &start)
	fmt.Printf("param start %v\n", &start)
}
