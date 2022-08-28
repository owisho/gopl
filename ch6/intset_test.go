package ch6

import (
	"fmt"
	"testing"
)

func TestIntset(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(&x)
	x.Remove(144)
	fmt.Println(&x)
	fmt.Println(x.Len())

	y.Add(9)
	y.Add(42)
	fmt.Println(&y)

	z := y.Copy()
	fmt.Println(z)

	x.UnionWith(&y)
	fmt.Println(x.String())
	fmt.Println(x.Has(9), x.Has(123))
}

func TestOperate(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)

	y.Add(9)
	y.Add(42)

	//x.UnionWith(&y)
	//x.IntersectWith(&y)
	x.DifferenceWith(&y)
	fmt.Println(&x)
}
