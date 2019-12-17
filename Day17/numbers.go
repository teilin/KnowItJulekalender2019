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

func leftRotate(num int) int {
	digits := numberOfDigits(num)
	pow10 := int(math.Pow10(digits - 1))
	firstDigit := num / pow10
	return ((num * 10) + firstDigit) - (firstDigit * pow10 * 10)
}

func calcTriangleNumber(n float64) float64 {
	return (n * (n + 1)) / 2
}

func isSquareNumber(n float64) bool {
	sr := math.Sqrt(n)
	return ((sr - math.Floor(sr)) == 0)
}

func isSquareNumberWithRotation(n float64, numTries int) {
	isSquare := isSquareNumber(n)

	if isSquare {
		countSquareNumbers += 1
	}

	if isSquare == false && numTries > 0 {
		i := int(n)
		i = leftRotate(i)
		isSquareNumberWithRotation(float64(i), numTries-1)
	}
}

func main() {
	for i := 0; i <= 1000000; i++ {
		triNum := calcTriangleNumber(float64(i))
		isSquareNumberWithRotation(triNum, numberOfDigits(int(triNum)))
	}
	fmt.Println(countSquareNumbers)
}
