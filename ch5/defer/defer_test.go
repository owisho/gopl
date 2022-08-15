package _defer

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	defer func() {
		fmt.Println("aaaaaa")
	}()
	defer trace("test")
	defer trace("test1")()
	fmt.Println("main")
}
