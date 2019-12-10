package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	dragonSize      int = 50
	leftoversheeps  int
	daysWithoutFood int
)

func readLines(path string) ([]int, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
	}
	inputTxt := string(data)

	tmp := strings.Split(inputTxt, ", ")
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

func dragonEat(n int) error {
	n += leftoversheeps
	leftoversheeps = 0
	if n >= dragonSize {
		leftoversheeps = n - dragonSize
		dragonSize += 1
		daysWithoutFood = 0
	} else if n < dragonSize {
		dragonSize -= 1
		daysWithoutFood++
	}
	if daysWithoutFood >= 5 {
		return errors.New("Dragon eats it all")
	}
	return nil
}

// KnowIt drageproblemer julekalender 1. desember 2019
func main() {
	data, err := readLines("sau.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	for index, input := range data {
		errDragonEat := dragonEat(input)
		if errDragonEat != nil {
			fmt.Println(errDragonEat)
			fmt.Printf("Days survived: %d\n", index)
		}
	}
}
