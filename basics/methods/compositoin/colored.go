package compositoin

import (
	"methods/geometry"
	"image/color"
)

type ColoredPoint struct {
	geometry.Point
	Color color.RGBA
}
