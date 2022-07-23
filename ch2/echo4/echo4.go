package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

	var k *Tmp = new(Tmp)
	p2 := k.A
	p1 := (*k).A
	fmt.Printf("equal %t\n", p1 == p2)
	var m Tmp
	p := m.A

	fmt.Println(p1, p, p2)

	var t = &Tmp{}
	fmt.Println(t)
	t.B = 10
	t.A = 1
	fmt.Println(t.B)
	fmt.Println(t.A)
	fmt.Println(t)

	var t1 Tmp
	fmt.Println(t1)
	t1.B = 10
	t1.A = 1
	fmt.Println(t1.B)
	fmt.Println(t1.A)
	fmt.Println(t1)

}

type Tmp struct {
	A int
	B int
}

func init() {

}
