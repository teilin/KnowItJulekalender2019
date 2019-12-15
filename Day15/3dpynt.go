package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	X, Y, Z float64
}

type Decor struct {
	Edge1, Edge2, Edge3 Coordinate
}

func string2Fload(number string) float64 {
	f, _ := strconv.ParseFloat(number, 32)
	return f
}

func readLines(path string) (map[int]Decor, error) {
	file, fileError := os.Open(path)
	if fileError != nil {
		fmt.Println("Error opening file")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	decors := make(map[int]Decor)
	index := 0
	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), ",")
		edge1 := &Coordinate{
			X: string2Fload(tmp[0]),
			Y: string2Fload(tmp[1]),
			Z: string2Fload(tmp[2]),
		}
		edge2 := &Coordinate{
			X: string2Fload(tmp[3]),
			Y: string2Fload(tmp[4]),
			Z: string2Fload(tmp[5]),
		}
		edge3 := &Coordinate{
			X: string2Fload(tmp[6]),
			Y: string2Fload(tmp[7]),
			Z: string2Fload(tmp[8]),
		}
		decor := Decor{
			Edge1: *edge1,
			Edge2: *edge2,
			Edge3: *edge3,
		}
		decors[index] = decor
		index += 1
	}
	return decors, fileError
}

func (v Coordinate) crossProduct(ov Coordinate) Coordinate {
	return Coordinate{
		X: v.Y*ov.Z - v.Z*ov.Y,
		Y: v.Z*ov.X - v.X*ov.Z,
		Z: v.X*ov.Y - v.Y*ov.X,
	}
}

func (coor1 Coordinate) dotProduct(coor2 Coordinate) float64 {
	return coor1.X*coor2.X + coor1.Y*coor2.Y + coor1.Z*coor2.Z
}

func massCalulation(decors map[int]Decor) float64 {
	var sum float64
	for _, value := range decors {
		sum += math.Abs(value.Edge1.crossProduct(value.Edge2).dotProduct(value.Edge3)) / 6
	}
	return sum
}

func main() {
	decorArray, readFileError := readLines("./EXAMPLE.csv")
	if readFileError != nil {
		fmt.Println("Read file error")
	}
	sum := massCalulation(decorArray)

	fmt.Println(sum)
	fmt.Println(sum / 1000.0)
}
