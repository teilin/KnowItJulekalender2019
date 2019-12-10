package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func FloatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func GetType(input string) string {
	if strings.Contains(input, "toalettpapir") {
		return "toalettpapir"
	}
	if strings.Contains(input, "sjampo") {
		return "sjampo"
	}
	if strings.Contains(input, "tannkrem") {
		return "tannkrem"
	}
	return "undefined"
}

func main() {
	var monthMap map[string]string = make(map[string]string)
	monthMap["Jan"] = "01"
	monthMap["Feb"] = "02"
	monthMap["Mar"] = "03"
	monthMap["Apr"] = "04"
	monthMap["May"] = "05"
	monthMap["Jun"] = "06"
	monthMap["Jul"] = "07"
	monthMap["Aug"] = "08"
	monthMap["Sep"] = "09"
	monthMap["Oct"] = "10"
	monthMap["Nov"] = "11"
	monthMap["Dec"] = "12"

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
			date := fmt.Sprintf("2018 %s %s 00 00", monthMap[dateInput[0:3]], dateInput[4:6])
			parsedDate, dateErr := time.Parse("2006 01 02 15 04", date)
			if dateErr != nil {
				fmt.Println("Error parseing date", dateErr)
			}
			weekDay = parsedDate.Weekday()
		} else {
			input := scanner.Text()
			inputType := GetType(input)
			i, _ := strconv.ParseFloat(re.FindString(scanner.Text()), 32)
			//fmt.Println(inputType + " => " + input + " => " + FloatToString(i))
			if inputType == "toalettpapir" {
				// Toalettpapir
				numToiletRolls += (i / 25)
				if weekDay == time.Wednesday {
					usedToiletPapirWednesday += i
				}
			}
			if inputType == "sjampo" {
				numShampooBottles += (i / 300.0)
				if weekDay == time.Sunday {
					usedShampooSunday += i
				}
			}
			if inputType == "tannkrem" {
				numTubesToothpaste += (i / 125)
			}
		}
		index += 1
	}

	fmt.Println(math.Ceil(numTubesToothpaste))
	fmt.Println(math.Ceil(numShampooBottles))
	fmt.Println(math.Ceil(numToiletRolls))
	fmt.Println(usedShampooSunday)
	fmt.Println(usedToiletPapirWednesday)

	fmt.Println(math.Ceil(numTubesToothpaste) * math.Ceil(numShampooBottles) * math.Ceil(numToiletRolls) * usedShampooSunday * usedToiletPapirWednesday)
}
