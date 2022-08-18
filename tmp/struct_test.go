package tmp

import (
	"fmt"
	"testing"
)

type A struct {
	Name string
	Age  int
}

type AA struct {
	*A
}

func TestTmp(t *testing.T) {
	aa := AA{
		A: &A{
			Name: "www",
			Age:  19,
		},
	}
	a := aa.A
	a.Name = "wy"
	bb := AA{
		a,
	}
	fmt.Println(bb.A)
	fmt.Println(aa.A)
	fmt.Println(a)
}
