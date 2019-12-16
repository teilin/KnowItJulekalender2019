package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	countCrossing int = 0
	fjord         [][]rune
	X, Y          int
	goNorth       bool = true
)

func sail() {
	for {
		if goNorth {
			if string(fjord[Y-3][X+1]) == "#" {
				goNorth = false
				countCrossing += 1
				X += 1
				Y += 1
			} else {
				X += 1
				Y -= 1
			}
		} else {
			if string(fjord[Y+3][X+3]) == "#" {
				goNorth = true
				countCrossing += 1
				X += 1
				Y -= 1
			} else {
				X += 1
				Y += 1
			}
		}
		if len(fjord[Y])-1 == X {
			break
		}
	}
}

func main() {
	file, fileError := os.Open("./fjord.txt")
	if fileError != nil {
		fmt.Println("Error opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fjord = append(fjord, []rune(scanner.Text()))
	}

	for index, _ := range fjord {
		for index2, value2 := range fjord[index] {
			if string(value2) == "B" {
				Y = index
				X = index2
			}
		}
	}

	sail()

	fmt.Println(countCrossing)
}
