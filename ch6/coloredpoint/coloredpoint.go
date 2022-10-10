package coloredpoint

import "image/color"
import "learn_go/ch6/geometry"

type ColoredPoint struct {
	geometry.Point
	Color color.RGBA
}
