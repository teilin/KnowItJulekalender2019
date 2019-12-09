package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readLines(path string) ([]string, error) {
	file, fileError := os.Open(path)
	if fileError != nil {
		fmt.Println("Error reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtLines []string
	for scanner.Scan() {
		txtLines = append(txtLines, scanner.Text())
	}

	return txtLines, fileError
}

func sumIntList(list []int64) int64 {
	var sum int64
	for _, input := range list {
		sum = sum + input
	}
	return sum
}

func krumpusSplit(n int64, number int64, splitIndex int) bool {
	numberStr := strconv.FormatInt(number, 10)

	if splitIndex > len(numberStr) {
		return false
	}

	a := numberStr[0:splitIndex]
	b := numberStr[splitIndex:len(numberStr)]

	if aNum, parseAError := strconv.ParseInt(a, 10, 64); parseAError == nil {
		if bNum, parseBError := strconv.ParseInt(b, 10, 64); parseBError == nil {
			if aNum == 0 || bNum == 0 {
				return krumpusSplit(n, number, splitIndex+1)
			}
			if aNum+bNum == n {
				fmt.Println("N = " + strconv.FormatInt(n, 10))
				fmt.Println("Number = " + strconv.FormatInt(number, 10))
				fmt.Println("A = " + strconv.FormatInt(aNum, 10))
				fmt.Println("B = " + strconv.FormatInt(bNum, 10))
				return true
			} else {
				if len(numberStr)-1 == splitIndex {
					return false
				}
				return krumpusSplit(n, number, splitIndex+1)
			}
		}
	}
	return krumpusSplit(n, number, splitIndex+1)
}

func isKrumpusTall(n string) (int64, bool) {
	var number int64
	var nsquared int64
	if v, parseErr := strconv.ParseInt(n, 10, 64); parseErr == nil {
		number = v
		nsquared = v * v
		isKrumpus := krumpusSplit(number, nsquared, 1)
		if isKrumpus {
			return number, isKrumpus
		}
	}
	return number, false
}

func main() {
	data, err := readLines("krampus.txt")

	if err != nil {
		fmt.Println("Error reading input file")
	}

	var krumpusNumbers []int64

	for _, input := range data {
		n, isKrumpusNumber := isKrumpusTall(input)
		if isKrumpusNumber {
			krumpusNumbers = append(krumpusNumbers, n)
		}
	}

	sumVal := sumIntList(krumpusNumbers)
	fmt.Println("Sum of list", sumVal)
}
