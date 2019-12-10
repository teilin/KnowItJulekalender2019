package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"time"
)

func FloatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func main() {
	var monthMap map[string]int = make(map[string]int)
	monthMap["Jan"] = 1
	monthMap["Feb"] = 2
	monthMap["Mar"] = 3
	monthMap["Apr"] = 4
	monthMap["May"] = 5
	monthMap["Jun"] = 6
	monthMap["Jul"] = 7
	monthMap["Aug"] = 8
	monthMap["Sep"] = 9
	monthMap["Oct"] = 10
	monthMap["Nov"] = 11
	monthMap["Dec"] = 12

	file, fileError := os.Open("logg.txt")
	if fileError != nil {
		fmt.Println("Error opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	var numTubesToothpaste float64 = 0
	var numShampooBottles float64 = 0
	var numToiletRolls float64 = 0
	var usedShampooSunday float64 = 0
	var usedToiletPapirWednesday float64 = 0
	var weekDay time.Weekday

	var index int = 0
	for scanner.Scan() {
		if index%4 == 0 {
			dateInput := scanner.Text()
			date := fmt.Sprintf("2018 %s %s 00 00", strconv.Itoa(monthMap[dateInput[0:3]]), dateInput[4:6])
			parsedDate, dateErr := time.Parse("2006 01 02 15 04", date)
			if dateErr != nil {
				fmt.Println("Error parseing date")
			}
			weekDay = parsedDate.Weekday()
			fmt.Println(weekDay)
		} else {
			if index%4 == 1 {
				// Tannkrem
				i, _ := strconv.ParseFloat(re.FindString(scanner.Text()), 32)
				numTubesToothpaste += (i / 125)
			}
			if index%4 == 2 {
				// Sjampoo
				i, _ := strconv.ParseFloat(re.FindString(scanner.Text()), 32)
				numShampooBottles += (i / 300.0)
				if weekDay == time.Sunday {
					usedShampooSunday += i
				}
			}
			if index%4 == 3 {
				// Toalettpapir
				i, _ := strconv.ParseFloat(re.FindString(scanner.Text()), 32)
				numToiletRolls += (i / 25)
				if weekDay == time.Wednesday {
					usedToiletPapirWednesday += i
				}
			}
		}
		index += 1
	}

	fmt.Println(math.Round(numTubesToothpaste))
	fmt.Println(math.Round(numShampooBottles))
	fmt.Println(math.Round(numToiletRolls))
	fmt.Println(usedShampooSunday)
	fmt.Println(usedToiletPapirWednesday)

	fmt.Println(math.Round(numTubesToothpaste) * math.Round(numShampooBottles) * math.Round(numToiletRolls) * usedShampooSunday * usedToiletPapirWednesday)
}
