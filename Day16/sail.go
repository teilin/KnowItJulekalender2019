package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	X, Y int
}

var (
	countCrossing int = 1
	fjord         [][]rune
	location      Point
	direction     Point = Point{X: 1, Y: -1}
	xlength       int   = 0
)

func elementExists(coordinate Point) bool {
	if coordinate.Y >= 0 && coordinate.Y < len(fjord) {
		if coordinate.X >= 0 && coordinate.X < len(fjord[coordinate.Y]) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func (current Point) CalcNextStep() Point {
	return Point{
		X: current.X + direction.X,
		Y: current.Y + direction.Y,
	}
}

func GetNextStep() {
	var threeStepAhead = location.CalcNextStep().CalcNextStep().CalcNextStep()
	if elementExists(threeStepAhead) {
		if string(fjord[threeStepAhead.Y][threeStepAhead.X]) == "#" {
			direction = Point{X: direction.X, Y: direction.Y * -1}
			countCrossing += 1
			location = Point{X: location.X + 1, Y: location.Y}
		} else {
			location = location.CalcNextStep()
		}
	} else {
		location = location.CalcNextStep()
	}
}

func sail() {
	for {
		GetNextStep()

		if location.X == xlength {
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

	xlength = len(fjord[0])

	for index, _ := range fjord {
		for index2, value2 := range fjord[index] {
			if string(value2) == "B" {
				location.X = index2
				location.Y = index
			}
		}
	}

	sail()
	fmt.Println(countCrossing)
}
