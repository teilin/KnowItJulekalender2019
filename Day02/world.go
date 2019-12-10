package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, fileError := os.Open("world.txt")
	if fileError != nil {
		fmt.Println("Error reading from file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	globalCount := 0
	for scanner.Scan() {
		current := scanner.Text()
		count := 0
		started := false
		prev := rune(0)
		for _, char := range current {
			if char == 32 && prev == 35 {
				started = true
			}
			if char == 35 && prev == 32 {
				started = false
				globalCount += count
				count = 0
			}
			if char == 32 && started {
				count += 1
			}
			prev = char
		}
	}

	fmt.Println(globalCount)
}
