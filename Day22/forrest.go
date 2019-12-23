package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput(path string) [][]rune {
	file, fileError := os.Open(path)
	if fileError != nil {
		fmt.Println("Error opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var data [][]rune
	for scanner.Scan() {
		data = append(data, []rune(scanner.Text()))
	}
	return data
}

func findRoot(forrest [][]rune) {
	/*x := len(forrest) - 1
	for y, _ := range forrest[x] {
		if forrest[y][x] == rune('#') {
			fmt.Println("")
			fmt.Printf("x %d y %d", x, y)
			height := getHeight(forrest, y, x)
			fmt.Println(height)
		}
	}*/
	for x, _ := range forrest[len(forrest)-1] {
		getHeight(forrest, len(forrest)-1, x)
	}
}

func getHeight(forrest [][]rune, xRoot int, yRoot int) int {
	height := 0
	for {
		height += 1
		yRoot -= 1
		if forrest[yRoot][xRoot] != rune('#') || yRoot <= 0 {
			break
		}
	}
	return height
}

func main() {
	forrest := readInput("./example.txt")
	findRoot(forrest)
}
