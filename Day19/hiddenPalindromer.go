package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sliceToInt(a []int64) int64 {
	if len(a) == 0 {
		return 0
	}

	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatInt(v, 10)
	}
	tmp, _ := strconv.ParseInt(strings.Join(b, ""), 10, 64)
	return tmp
}

func intToSlice(num int64) []int64 {
	if num == 0 {
		return []int64{0}
	}
	var slice []int64
	for {
		slice = append(slice, num%10)
		num /= 10
		if num <= 0 {
			break
		}
	}
	return reverseSlice(slice)
}

func sumSlice(slice []int64) int64 {
	var sum int64 = 0
	for _, i := range slice {
		sum += i
	}
	return sum
}

func reverseSlice(s []int64) []int64 {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func isPalidrome(num int64) bool {
	rs := true
	numSlice := intToSlice(num)
	numSliceReversed := reverseSlice(numSlice)
	for i, j := 0, len(numSliceReversed)-1; i <= j; i, j = i+1, j-1 {
		rs = numSlice[i] == numSliceReversed[j]
		if !rs {
			break
		}
	}
	return rs
}

func isHiddenPalindromer(num int64) bool {
	if isPalidrome(num) {
		return false
	}
	n := num + sliceToInt(reverseSlice(intToSlice(num)))
	return isPalidrome(n)
}

func main() {
	var sum int64 = 0
	var i int64
	for i = 1; i < 123454321; i++ {
		if isHiddenPalindromer(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}
