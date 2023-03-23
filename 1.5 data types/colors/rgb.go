package colors

import (
	"fmt"
	"math"
)

type RGBColor struct {
	R uint8
	G uint8
	B uint8
}

func (c RGBColor) ToText() string {
	return fmt.Sprintf("RGB { %v, %v, %v }", c.R, c.G, c.B)
}

func (c RGBColor) ToHEX() string {
	return fmt.Sprintf("#%X%X%X", c.R, c.G, c.B)
}

func (c RGBColor) Luminance() uint8 {
	return uint8(math.Round(0.2126*float64(c.R) + 0.7152*float64(c.G) + 0.0722*float64(c.B)))
}

// Init can be used as Constructor for validate the data before create the object instance
func InitRGBColor(R uint8, G uint8, B uint8) RGBColor {
	return RGBColor{R, G, B}
}
