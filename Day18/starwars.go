package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Employee struct {
	firstName string
	lastName  string
	sex       string
}

var (
	manNames      []string
	womanNames    []string
	part1LastName []string
	part2LastName []string
	alphabeth     map[string]int = make(map[string]int)
	employees     []Employee
)

func init() {
	alphabeth["a"] = 1
	alphabeth["b"] = 2
	alphabeth["c"] = 3
	alphabeth["d"] = 4
	alphabeth["e"] = 5
	alphabeth["f"] = 6
	alphabeth["g"] = 7
	alphabeth["h"] = 8
	alphabeth["i"] = 9
	alphabeth["j"] = 10
	alphabeth["k"] = 11
	alphabeth["l"] = 12
	alphabeth["m"] = 13
	alphabeth["n"] = 14
	alphabeth["o"] = 15
	alphabeth["p"] = 16
	alphabeth["q"] = 17
	alphabeth["r"] = 18
	alphabeth["s"] = 19
	alphabeth["t"] = 20
	alphabeth["u"] = 21
	alphabeth["v"] = 22
	alphabeth["w"] = 23
	alphabeth["x"] = 24
	alphabeth["y"] = 25
	alphabeth["z"] = 26
	alphabeth["æ"] = 27
	alphabeth["ø"] = 28
	alphabeth["å"] = 29
}

func intToSlice(num int) []int {
	if num == 0 {
		return []int{0}
	}
	var slice []int
	for {
		slice = append(slice, num%10)
		num /= 10
		if num <= 0 {
			break
		}
	}
	return slice
}

func SplitToInt(a []int, sep string) int {
	if len(a) == 0 {
		return 0
	}

	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}
	tmp, _ := strconv.Atoi(strings.Join(b, sep))
	return tmp
}

func sumASCII(word string) int {
	runes := []rune(word)
	sum := 0
	for _, char := range runes {
		sum += int(char)
	}
	return sum
}

func multiplyASCII(word string) int {
	runes := []rune(word)
	sum := 1
	for _, char := range runes {
		sum *= int(char)
	}
	return sum
}

func sumPositionAlphabeth(word string) int {
	runes := []rune(word)
	sum := 0
	for _, char := range runes {
		t := strings.ToLower(string(char))
		sum += alphabeth[t]
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
		e := Employee{firstName: employee[0], lastName: employee[1], sex: employee[2]}
		employees = append(employees, e)
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
			}
			tmp = tmp[:0]
			index += 1
		} else {
			tmp = append(tmp, line)
		}
	}
	for _, v := range tmp {
		part2LastName = append(part2LastName, v)
	}
}

func firstName2StarWarsName(emp Employee) string {
	if emp.sex == "M" {
		count := len(manNames)
		return manNames[sumASCII(emp.firstName)%count]
	} else {
		count := len(womanNames)
		return womanNames[sumASCII(emp.firstName)%count]
	}
}

func lastName2StarWarsName(employee Employee) (string, string) {
	lengthLastName := len(employee.lastName)
	lastPart1 := employee.lastName[0:int(math.Ceil(float64(lengthLastName)/2))]
	lastPart2 := employee.lastName[int(math.Ceil(float64(lengthLastName)/2)):lengthLastName]

	var starWarsLastName1 string
	var starWarsLastName2 string

	starWarsLastName1 = part1LastName[sumPositionAlphabeth(lastPart1)%len(part1LastName)]

	if employee.sex == "M" {
		count := len(employee.firstName)
		index := multiplyASCII(lastPart2) * count
		slice := intToSlice(index)
		sort.Sort(sort.Reverse(sort.IntSlice(slice)))
		starWarsLastName2 = part2LastName[SplitToInt(slice, "")%len(part2LastName)]
	} else {
		count := len(employee.firstName) + len(employee.lastName)
		index := multiplyASCII(lastPart2) * count
		slice := intToSlice(index)
		sort.Sort(sort.Reverse(sort.IntSlice(slice)))
		starWarsLastName2 = part2LastName[SplitToInt(slice, "")%len(part2LastName)]
	}

	return starWarsLastName1, starWarsLastName2
}

func main() {
	readEmployees("./employees.csv")
	readStarWarsNames("./names.txt")

	var starWarsNames map[string]int = make(map[string]int)
	var mostUsedName string
	var mostUsedNameCount int = 0

	for _, employee := range employees {
		starWarsFirstName := firstName2StarWarsName(employee)
		starWarsLastName1, starWarsLastName2 := lastName2StarWarsName(employee)

		name := fmt.Sprintf("%s %s%s", starWarsFirstName, starWarsLastName1, starWarsLastName2)
		starWarsNames[name]++
		if starWarsNames[name] > mostUsedNameCount {
			mostUsedNameCount = starWarsNames[name]
			mostUsedName = name
		}
	}
	fmt.Println(mostUsedName)
}
