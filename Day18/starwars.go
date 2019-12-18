package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

var (
	manNames      []string
	womanNames    []string
	part1LastName []string
	part2LastName []string
)

func sumASCII(word string) int {
	runes := []rune(word)
	sum := 0
	for _, char := range runes {
		sum += int(char)
	}
	return sum
}

func readEmployees(path string) {
	file, fileError := os.Open(path)
	if fileError != nil {
		fmt.Println("Error opening file")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, csvErr := reader.ReadAll()
	if csvErr != nil {
		fmt.Println("Error reading csv file")
	}
	for _, employee := range records {
		fmt.Println(employee[0])
	}
}

func readStarWarsNames(path string) {
	file, fileError := os.Open(path)
	if fileError != nil {
		fmt.Println("Error opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var tmp []string
	index := 1

	for scanner.Scan() {
		line := scanner.Text()
		if line == "---" {
			if index == 1 {
				for _, v := range tmp {
					manNames = append(manNames, v)
				}
			} else if index == 2 {
				for _, v := range tmp {
					womanNames = append(womanNames, v)
				}
			} else if index == 3 {
				for _, v := range tmp {
					part1LastName = append(part1LastName, v)
				}
			} else if index == 4 {
				for _, v := range tmp {
					part2LastName = append(part2LastName, v)
				}
			}
			tmp = tmp[:0]
			index += 1
		} else {
			tmp = append(tmp, line)
		}
	}

	fmt.Println(part2LastName)
}

func main() {
	//readEmployees("./employees.csv")
	readStarWarsNames("./names.txt")
}
