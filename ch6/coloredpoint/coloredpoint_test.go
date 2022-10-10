package coloredpoint

import (
	"fmt"
	"learn_go/ch6/geometry"
	"testing"
)
import "image/color"

func TestColoredPoint(t *testing.T) {
	p := ColoredPoint{Point: geometry.Point{X: 10, Y: 10}, Color: color.RGBA{A: 0, G: 0, B: 0, R: 0}}
	p.X = 1
	p.Y = 2
	p.ScaleBy(2)
	fmt.Println(p)
}
