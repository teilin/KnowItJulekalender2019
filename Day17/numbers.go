package main

import (
	"fmt"
	"math"
)

var (
	countSquareNumbers int = 0
)

func numberOfDigits(n int) int {
	cnt := 0
	for {
		cnt += 1
		n /= 10
		if n <= 0 {
			break
		}
	}
	return cnt
}

func leftRotate(num int) []int {
	digits := numberOfDigits(num)
	var iSlice []int = make([]int, digits)
	pow10 := int(math.Pow10(digits - 1))
	for i := 0; i < digits; i++ {
		firstDigit := num / pow10
		left := ((num * 10) + firstDigit) - (firstDigit * pow10 * 10)
		num = left
		iSlice[i] = left
	}
	return iSlice
}

func calcTriangleNumber(n int) int {
	return (n * (n + 1)) / 2
}

func isSquareNumber(n float64) bool {
	sr := math.Sqrt(n)
	return ((sr - math.Floor(sr)) == 0)
}

func isSquareNumberWithRotation(n float64) {
	isSquare := isSquareNumber(n)

	if isSquare {
		countSquareNumbers += 1
	}

	if isSquare == false {
		i := int(n)
		iLeft := leftRotate(i)
		for _, v := range iLeft {
			isSquare = isSquareNumber(float64(v))
			if isSquare {
				countSquareNumbers += 1
				break
			}
		}
	}
}

func main() {
	for i := 0; i <= 1000000; i++ {
		triNum := calcTriangleNumber(i)
		isSquareNumberWithRotation(float64(triNum))
	}
	fmt.Println(countSquareNumbers)
}
