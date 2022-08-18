package tmp

import (
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	f1()
}

func f1() {
	defer func() {
		fmt.Println("defer1")
		//panic("***** defer1 *****")
	}()
	defer func() {
		fmt.Println("defer2")
		panic("------ defer2 ------")
	}()
}
