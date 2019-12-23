package main

import (
	"fmt"
	"math/big"
)

func sumIntDigits(num int64) int64 {
	if num == 0 {
		return 0
	}
	var slice []int64
	for {
		slice = append(slice, num%10)
		num /= 10
		if num <= 0 {
			break
		}
	}
	var sum int64
	for _, v := range slice {
		sum += v
	}
	return sum
}

func main() {
	count := 0
	for i := 1; i <= 98765432; i++ {
		sumDigits := sumIntDigits(int64(i))
		isHarshadNumber := int64(i)%sumDigits == 0
		if isHarshadNumber {
			if big.NewInt(sumDigits).ProbablyPrime(0) {
				count += 1
			}
		}
	}
	fmt.Println(count)
}
