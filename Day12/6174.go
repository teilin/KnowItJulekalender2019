package main

import (
	"fmt"
	"sort"
	"strconv"
)

func sliceToInt(s []int) int {
	res := 0
	op := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] * op
		op *= 10
	}
	return res
}

func stringToIntSlice(number string) []int {
	if _, err := strconv.Atoi(number); err != nil {
		fmt.Println("Error parseing")
	}
	var slice []int
	for i := 0; i <= (4 - len(number) - 1); i++ {
		slice = append(slice, 0)
	}
	for _, digit := range number {
		slice = append(slice, int(digit)-int('0'))
	}
	return slice
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func calc(number int) int {
	str := strconv.Itoa(number)
	decSlice := stringToIntSlice(str)
	incSlice := stringToIntSlice(str)
	sort.Sort(sort.Reverse(sort.IntSlice(decSlice)))
	sort.Sort(sort.IntSlice(incSlice))
	dec := sliceToInt(decSlice)
	inc := sliceToInt(incSlice)
	return Max(dec, inc) - Min(dec, inc)
}

func repeatToNumber(number int, index int) (int, int) {
	if number == 6174 {
		return number, index
	}
	if index > 7 {
		return number, -1
	}
	num := calc(number)
	return repeatToNumber(num, index+1)
}

func main() {
	count := 0

	for i := 1000; i <= 9999; i++ {
		_, itteration := repeatToNumber(i, 0)
		if itteration == 7 {
			count += 1
		}
	}

	fmt.Println(count)
}
