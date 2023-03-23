package main

import (
	"fmt"
	"workspace/colors"
)

func main() {
	var colorR uint8 = 255

	mycolor := colors.InitRGBColor(colorR, 255, 255)

	fmt.Println("my color is: ", mycolor.ToText(), " => ", mycolor.ToHEX())
	fmt.Printf("the color Luminance is %v\n", mycolor.Luminance())
}
