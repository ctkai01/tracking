package utils

import (
	"image/color"
	"strconv"
	"time"
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



func ConvertSecToHourMinSec(seconds int) (hour, minute, second string) {
	duration := time.Duration(seconds) * time.Second
	hour = formatNumber2String(int(duration.Hours()), HOUR)
	minute = formatNumber2String(int(duration.Minutes()) % 60, MINUTE)
	second = formatNumber2String(int(duration.Seconds()) % 60, SECOND)
	return
}

func formatNumber2String(num int, timeType int) string {
	strNum := strconv.Itoa(num)

	if num < 10 {
		strNum = "0" + strNum
	}

	// if timeType == HOUR {
	// 	strNum = strNum + ":"
	// }

	// if timeType == MINUTE {
	// 	strNum = strNum + ":"
	// }

	return strNum
}