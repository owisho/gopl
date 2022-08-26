package tmp

import (
	"fmt"
	"learn_go/ch7"
	"math"
	"testing"
	"time"
)

type B int

//func (b *B) toString() string {
//	fmt.Printf("b=%v\n", b)
//	*b = 3
//	return fmt.Sprintf("%d", *b)
//}

func (b B) toString() string {
	fmt.Printf("b=%v\n", b)
	b = 3
	return fmt.Sprintf("%d", b)
}

func TestFunction(t *testing.T) {
	var b B = 2
	fmt.Println((&b).toString())
	fmt.Println(b)
}

type Tmp struct {
	*Point
	ch7.LineCounter
}

func TestTt(t *testing.T) {
	b := Tmp{
		&Point{X: 1, Y: 2},
		1,
	}
	fmt.Println(b.Point)
	time.AfterFunc(10*time.Second, pp)
}

func pp() {
	fmt.Println("11")
}

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X = factor
	p.Y = factor
}

func TestScale(t *testing.T) {
	p := Point{X: 1, Y: 2}
	fmt.Println(p)
	fmt.Printf("distance=%f\n", Point{X: 1, Y: 2}.Distance(p))
	(&p).ScaleBy(2)
	fmt.Println(p)
}

func TestMapLen(t *testing.T) {
	m := map[string]string{
		"key1": "e",
		"key2": "ee",
	}
	fmt.Println(len(m))
}

type Values map[string][]string

func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}

func returnNilMap() map[string]string {
	return nil
}

func returnNilSlice() []string {
	return nil
}

func TestNil(t *testing.T) {
	var v Values = make(map[string][]string)
	v.Add("aa", "bbb")
	fmt.Println(v)
}

type FF struct {
	AA string
	BB string
}

func TestNil1(t *testing.T) {
	var f FF
	fmt.Println(f)
}
