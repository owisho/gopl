package geometry

import (
	"fmt"
	"testing"
	"time"
)

func TestDistance(t *testing.T) {
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())
}

func TestMethod(t *testing.T) {
	p := Point{1, 2}
	q := Point{4, 6}

	distanceFromP := p.Distance
	fmt.Println(distanceFromP(q))
	var origin Point
	fmt.Println(distanceFromP(origin))

	scaleP := p.ScaleBy
	scaleP(2)
	scaleP(3)
	scaleP(10)
}

type Rocket struct {
}

func (r *Rocket) Launch() {}

func TestMethod1(t *testing.T) {
	r := new(Rocket)
	time.AfterFunc(10*time.Second, r.Launch)
}
