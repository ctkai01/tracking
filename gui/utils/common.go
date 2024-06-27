package utils

import (
	"image/color"
	"strconv"
)

func TransformColorFromHex(hexColor string) color.NRGBA {
	hexColor = hexColor[1:]

	// Parse the hexadecimal components (R, G, B, and A)
	red, _ := strconv.ParseInt(hexColor[0:2], 16, 64)
	green, _ := strconv.ParseInt(hexColor[2:4], 16, 64)
	blue, _ := strconv.ParseInt(hexColor[4:6], 16, 64)
	alpha := uint8(255) // Assume full opacity (alpha = 255)

	// Create the NRGBA color
	nrgbaColor := color.NRGBA{
		R: uint8(red),
		G: uint8(green),
		B: uint8(blue),
		A: alpha,
	}

	return nrgbaColor
}

func GetDigits(number int, timeCounterSize int) []uint {
	digits := make([]uint, timeCounterSize)
	for i := 0; i < timeCounterSize; i++ {
		digits[timeCounterSize-i-1] = uint(number % 10)
		number /= 10
	}

	return digits
}
