package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sledgespeed      int = 10703437
	prevChar         string
	numIcePairs      int
	numMountainPairs int
	numTerrains      int
)

func calc(input string) int {
	var terrain map[string]int = make(map[string]int)
	terrain["G"] = -27
	terrain["I"] = 12
	terrain["A"] = -59
	terrain["S"] = -212
	terrain["F"] = -70

	numTerrains += 1

	if input == "I" {
		numIcePairs += 1
	}

	if input == "I" {
		sledgespeed += terrain[input] * numIcePairs
	} else {
		if prevChar == "F" {
			sledgespeed += 35
		} else {
			sledgespeed += terrain[input]
		}
	}

	if prevChar == "I" && input != "I" {
		numIcePairs = 0
	}

	if prevChar == "F" {
		prevChar = "FF"
	} else {
		prevChar = input
	}
	return numTerrains
}

func main() {
	file, fileError := os.Open("./terreng.txt")
	if fileError != nil {
		fmt.Println("Error in opening file")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		if sledgespeed > 0 {
			calc(scanner.Text())
		}
	}

	fmt.Println("Sledge stops after " + strconv.Itoa(numTerrains) + "km")
}
