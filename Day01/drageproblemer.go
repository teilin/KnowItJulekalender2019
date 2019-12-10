package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readLines(path string) ([]int, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
	}
	inputTxt := string(data)

	tmp := strings.Split(inputTxt, ",")
	values := make([]int, 0, len(tmp))
	for _, raw := range tmp {
		v, err := strconv.Atoi(raw)
		if err != nil {
			fmt.Println("Error parsing string to int", raw)
			continue
		}
		values = append(values, v)
	}
	return values, err
}

// KnowIt drageproblemer julekalender 1. desember 2019
func main() {
	data, err := readLines("sau.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	for _, input := range data {
		fmt.Println("Element", input)
	}
}
