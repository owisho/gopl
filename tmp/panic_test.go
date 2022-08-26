package tmp

import (
	"fmt"
	"os"
	"runtime"
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

func TestF(t *testing.T) {
	defer printStack()
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
